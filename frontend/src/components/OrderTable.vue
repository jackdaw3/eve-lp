<template>
  <div class="OrderTable">
    <el-table :data="orderTableData" :cell-style="tableFormat" style="width: 100%" v-loading="loading"
      :element-loading-spinner="SVG" element-loading-svg-view-box="-10, -10, 50, 50" :table-layout="tableLayout"
      :highlight-current-row="highlight">
      <el-table-column prop="OrderId" :label="t('order.orderId')" min-width="15%"></el-table-column>
      <el-table-column prop="SystemName" :label="t('order.systemName')" min-width="15%"></el-table-column>
      <el-table-column :label="t('order.volume')" :formatter="volumeFormat" min-width="15%"></el-table-column>
      <el-table-column prop="Price" :label="t('order.price')" :formatter="numberFormat" sortable
        min-width="18%"></el-table-column>
      <el-table-column prop="Expiration" :label="t('order.expiration')" min-width="22%"></el-table-column>
      <el-table-column prop="LastUpdated" :label="t('order.lastUpdated')" min-width="15%"></el-table-column>
    </el-table>
  </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, reactive, ref, watch } from 'vue';
import { useI18n } from 'vue-i18n';
import { useStore } from 'vuex';
import { ElTable, ElTableColumn } from 'element-plus';
import axios from 'axios';
import { BACKEND_SERVER, SVG } from '~/constants';
interface Order {
  OrderId: string;
  SystemName: string;
  VolumeRemain: number;
  VolumeTotal: number;
  Price: number;
  Expiration: string;
  LastUpdated: string;
}

const { t, locale } = useI18n();
const store = useStore();
const props = defineProps<{
  orderData: {
    regionId: string;
    offerId: string;
    materialPrice: string;
    weightedPrice: string;
    tax: string;
  };
  buyOrder: boolean
}>();
const orderTableData = reactive<Order[]>([])
let loading = ref(false)
const tableLayout = computed(() => store.state.tableForm.layout);
const highlight = computed(() => store.state.tableForm.highlight === "true");

onMounted(() => {
  getOrders()
})

watch(locale, () => {
  getOrders()
})

function numberFormat(row: any, column: any, cellValue: number) {
  cellValue = Math.floor(cellValue);
  let cellValueStr = cellValue.toString();
  if (!cellValueStr.includes('.')) cellValueStr += '.';
  return cellValueStr
    .replace(/(\d)(?=(\d{3})+\.)/g, '$1,')
    .replace(/\.$/, '');
}

function volumeFormat(row: Order) {
  return `${row.VolumeRemain}/${row.VolumeTotal}`;
}

function tableFormat() {
  return {
    padding: '0',
    fontSize: '14px',
  };
}

async function getOrders() {
  loading.value = true
  try {
    const response = await axios.get(BACKEND_SERVER + "order", {
      params: {
        regionId: props.orderData.regionId,
        offerId: props.orderData.offerId,
        scope: props.orderData.weightedPrice,
        tax: props.orderData.tax,
        isBuyOrder: props.buyOrder,
        lang: locale.value,
      },
    });

    orderTableData.splice(0, orderTableData.length, ...response.data);
  } catch (error) {
    console.error(error);
  } finally {
    loading.value = false
  }
}

</script>

<style scoped></style>