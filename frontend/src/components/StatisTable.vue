<template>
  <div class="StatisTable">
    <el-table :data="statisTableData" :cell-style="tableStyle" :table-layout="tableLayout"
      :highlight-current-row="highlight" stle="width: 100%" v-loading="loading" :element-loading-spinner="SVG"
      element-loading-svg-view-box="-10, -10, 50, 50">
      <el-table-column type="expand">
        <template #default="props">
          <el-table :data="props.row.Orderwrappers" :cell-style="tableStyle" table-layout="auto" :header-cell-style="{
            padding: '0',
          }">
            <el-table-column prop="OrderDTO.SystemName" :label="t('order.systemName')"
              min-width="15%"></el-table-column>
            <el-table-column :label="t('order.volume')" :formatter="volumeFormat" min-width="15%"></el-table-column>
            <el-table-column prop="OrderDTO.Price" :label="t('order.price')" :formatter="numberFormat" sortable
              min-width="20%"></el-table-column>
            <el-table-column prop="Income" :label="t('order.statis.income')" :formatter="numberFormat" sortable
              min-width="20%"></el-table-column>
            <el-table-column prop="Cost" :label="t('order.statis.cost')" :formatter="numberFormat" sortable
              min-width="20%"></el-table-column>
            <el-table-column prop="Profit" :label="t('order.statis.profit')" :formatter="numberFormat" sortable
              min-width="20%"></el-table-column>
            <el-table-column prop="UnitProfit" :label="t('order.statis.unitProfit')" sortable
              min-width="15%"></el-table-column>
          </el-table>
        </template>
      </el-table-column>
      <el-table-column prop="UnitProfitRange" :label="t('order.statis.lpRange')" min-width="25%"></el-table-column>
      <el-table-column prop="Quantity" :label="t('order.statis.number')" min-width="11%"></el-table-column>
      <el-table-column prop="Income" :label="t('order.statis.income')" :formatter="numberFormat" sortable
        min-width="16%"></el-table-column>
      <el-table-column prop="Cost" :label="t('order.statis.cost')" :formatter="numberFormat" sortable
        min-width="16%"></el-table-column>
      <el-table-column prop="Profit" :label="t('order.statis.profit')" :formatter="numberFormat" sortable
        min-width="16%"></el-table-column>
      <el-table-column prop="AveUnitProfit" :label="t('order.statis.avgLpPrice')" sortable
        min-width="16%"></el-table-column>
    </el-table>
  </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, reactive, ref, watch } from 'vue';
import { useI18n } from 'vue-i18n';
import { useStore } from 'vuex';
import { ElTable, ElTableColumn } from 'element-plus';
import { BACKEND_SERVER, SVG } from '~/constants';
import axios from 'axios';

interface DataItem {
  Orderwrappers: OrderWrapper[];
  UnitProfitRange: string;
  Quantity: number;
  Income: number;
  Cost: number;
  Profit: number;
  AveUnitProfit: number;
}

interface OrderWrapper {
  OrderDTO: {
    SystemName: string;
    Price: number;
    VolumeRemain: number;
    VolumeTotal: number;
  };
  Income: number;
  Cost: number;
  Profit: number;
  UnitProfit: number;
}

const { t, locale } = useI18n();
const store = useStore();
const props = defineProps<{
  statisData: {
    regionId: string;
    offerId: string;
    materialPrice: string;
    weightedPrice: string;
    tax: string;
  };
  buyOrder: boolean
}>();
let statisTableData = reactive<DataItem[]>([])
let loading = ref(false)
const tableLayout = computed(() => store.state.tableForm.layout);
const highlight = computed(() => store.state.tableForm.highlight === "true");

onMounted(() => {
  getStatis()
})

watch(locale, () => {
  getStatis()
})

function numberFormat(row: any, column: any, cellValue: number) {
  cellValue = Math.floor(cellValue);
  let cellValueStr = cellValue.toString();
  if (!cellValueStr.includes('.')) cellValueStr += '.';
  return cellValueStr
    .replace(/(\d)(?=(\d{3})+\.)/g, '$1,')
    .replace(/\.$/, '');
}

function tableStyle() {
  return {
    padding: '0',
    fontSize: '14px',
  };
}

function volumeFormat(row: OrderWrapper) {
  return `${row.OrderDTO.VolumeRemain}/${row.OrderDTO.VolumeTotal}`;
}

async function getStatis() {
  loading.value = true
  try {
    const response = await axios.get(BACKEND_SERVER + "statis", {
      params: {
        offerId: props.statisData.offerId,
        regionId: props.statisData.regionId,
        scope: props.statisData.weightedPrice,
        materialPrice: props.statisData.materialPrice,
        tax: props.statisData.tax,
        isBuyOrder: props.buyOrder,
        lang: locale.value,
      },
    });

    statisTableData.splice(0, statisTableData.length, ...response.data);
  } catch (error) {
    console.error(error);
  } finally {
    loading.value = false
  }
}

</script>

<style scoped></style>