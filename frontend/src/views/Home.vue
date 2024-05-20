<template>
  <div>
    <el-config-provider :locale="elLocale">
      <Header :region-visible="true" />
      <div>
        <div :style="{ marginTop: getMarginTop() + 'px', marginLeft: '15px'}">
          <Config @exportRequest="handleExportExcelRequest" />
        </div>
        <Table class="table" ref="tableRef" />
        <Item />
        <Calculator />
      </div>
      <el-backtop :right="50" :bottom="50" />
    </el-config-provider>
  </div>
</template>

<script lang='ts' setup>
import { computed, ref } from 'vue';
import { useI18n } from 'vue-i18n';
import { ElConfigProvider } from 'element-plus';
import { elementPlusLocales } from '~/i18n/i18n';
import Item from '~/components/Item.vue';

const { locale } = useI18n();
const elLocale = computed(() => { 
  return elementPlusLocales[locale.value] || elementPlusLocales.en; 
});
const tableRef = ref<any>()

function handleExportExcelRequest(corporationName: string) {
  if (tableRef.value) {
    tableRef.value.exportExcel(corporationName)
  }
}

function getMarginTop() {
  return document.documentElement.clientHeight * 0.07
}
</script>

<style scoped>
#app {
  text-align: center;
  color: var(--el-text-color-primary);
}

.config {
  margin-left: 10px;
}

.table {
  display: flex;
  flex-direction: column;
  align-items: center;
}
</style>
