<template>
    <div style="margin-top: 0.3%">
        <el-table :data="filterTableData" :table-layout="tableLayout" :highlight-current-row="highlight"
            style="width: 98.5%;" :cell-style="tableStyle" v-loading="loading" :element-loading-spinner="SVG"
            element-loading-svg-view-box="-10, -10, 50, 50" :header-cell-style="{ height: '60px' }"
            :row-class-name="handelRowDetail" @sort-change="sortChange" @selection-change="handleSelectionChange"
            id="table">
            <el-table-column type="expand">
                <template #default="props">
                    <el-table :data="props.row.Matertials" style="width: 55%" :table-layout="tableLayout"
                        :cell-style="tableStyle" :span-method="objectSpanMethod" :row-class-name="handelRowDetail"
                        :header-cell-style="{ padding: '0' }">
                        <el-table-column prop="IsBluePrint" align="center" min-width="12%">
                            <template #default="scope">
                                <span v-if="scope.row.IsBluePrint === true">
                                    {{ t('table.material.bluePrintMaterial') }}
                                </span>
                                <span v-else>{{ t('table.material.lpStoreMaterial') }}</span>
                            </template>
                        </el-table-column>
                        <el-table-column prop="MaterialName" :label="t('table.material.name')" min-width="30%">
                            <template #default="scope">
                                <el-image style="height: 22px; vertical-align: middle" :src="getIcon(scope.row.ItemId)"
                                    @click="openItem(scope.row)" fit="contain" class="grab-cursor" lazy>
                                    <template #placeholder>
                                        <div style="height: 22px; background-color: transparent;"></div>
                                    </template>
                                    <template #error>
                                        <div class="image-slot">
                                            <el-icon><icon-picture /></el-icon>
                                        </div>
                                    </template>
                                </el-image>
                                <span>{{ scope.row.MaterialName }}</span>
                            </template>
                        </el-table-column>
                        <el-table-column prop="Quantity" :label="t('table.material.quantity')"
                            :formatter="numberFormat" min-width="10%" />
                        <el-table-column prop="Price" :label="t('table.material.price')" :formatter="numberFormat"
                            min-width="14%" />
                        <el-table-column prop="Cost" :label="t('table.material.cost')" :formatter="numberFormat"
                            min-width="14%" />
                        <el-table-column min-width="20%">
                            <template #header>
                                <el-button size="small" @click="copyAllMaterials(props.row.Matertials)" plain>{{
                                    t('table.material.copy') }}</el-button>
                            </template>
                            <template #default="scope">
                                <el-button size="small" @click="copyMaterial(scope.row)" plain>{{
                                    t('table.material.copy') }}</el-button>
                                <el-button v-if="scope.row.Error === true" size="small" plain
                                    @click="errorMessage(scope.row.ErrorMessage)" style="margin-left: 0px">{{
                                        t('table.material.error')
                                    }}</el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                </template>
            </el-table-column>
            <el-table-column type="selection" v-if="calculator" />
            <el-table-column prop="ItemName" :label="t('table.name')" min-width="20%">
                <template #default="scope">
                    <el-image style="height: 22px; vertical-align: middle" :src="getIcon(scope.row.ItemId)"
                        @click="openItem(scope.row)" fit="contain" class="grab-cursor" lazy>
                        <template #placeholder>
                            <div style="height: 22px; background-color: transparent;"></div>
                        </template>
                        <template #error>
                            <div class="image-slot">
                                <el-icon><icon-picture /></el-icon>
                            </div>
                        </template>
                    </el-image>
                    <span>{{ scope.row.ItemName }}</span>
                </template>
            </el-table-column>
            <el-table-column prop="Quantity" :label="t('table.quantity')" :formatter="numberFormat" min-width="5%" />
            <el-table-column prop="LpCost" :label="t('table.lpCost')" :formatter="numberFormat" sortable="custom"
                min-width="7%" />
            <el-table-column prop="IskCost" :label="t('table.iskCost')" :formatter="numberFormat" sortable="custom"
                min-width="7%" />
            <el-table-column prop="MaterialCost" :label="t('table.materialCost')" :formatter="numberFormat"
                sortable="custom" min-width="8%" />
            <el-table-column prop="Price" :label="t('table.price')" :formatter="numberFormat" sortable="custom"
                min-width="8%" />
            <el-table-column prop="Income" :label="t('table.income')" :formatter="numberFormat" sortable="custom"
                min-width="8%" />
            <el-table-column prop="Profit" :label="t('table.profit')" :formatter="numberFormat" sortable="custom"
                min-width="8%" />
            <el-table-column prop="Volume" :label="t('table.volume')" :formatter="numberFormat" sortable="custom"
                min-width="6%" />
            <el-table-column prop="SaleIndex" :label="t('table.saleIndex')" :formatter="numberFormat" sortable="custom"
                min-width="7%" />
            <el-table-column prop="UnitProfit" :label="t('table.unitProfit')" sortable="custom" min-width="6%" />
            <el-table-column :width="operationWidth" min-width="10%">
                <template #header align="right">
                    <el-input v-model="search" size="default" :placeholder="t('table.lookUp')" :prefix-icon="Search" />
                </template>
                <template #default="scope">
                    <el-button size="small" plain @click="openOrder(scope)">{{ t('table.order') }}</el-button>
                    <el-button v-if="scope.row.Error === true" size="small" plain
                        @click="errorMessage(scope.row.ErrorMessage)" style="margin-left: 0px">{{ t('table.error')
                        }}</el-button>
                </template>
            </el-table-column>
            <template #empty>
                <div class="custom-empty">
                    <el-empty :image-size="250" />
                </div>
            </template>
        </el-table>
        <br />

        <div style="display: flex; align-items: center; justify-content: space-between;width: 100%">
            <div style="flex: 0;">
                <el-button v-if="calculator" :icon="TakeawayBox" size="large" circle style="margin-left: 15px;"
                    @click="openCalculator()"></el-button>
            </div>
            <div style="flex: 1; display: flex; justify-content: center;">
                <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                    :current-page="currentPage" :page-sizes="[25, 50, 100, 200, 500]" :page-size="pageSize"
                    layout="total, sizes, prev, pager, next, jumper" :total="tableData.length">
                </el-pagination>
            </div>
        </div>

    </div>
</template>

<script lang="ts" setup>
import axios from 'axios';
import { computed, nextTick, onMounted, reactive, ref, watch } from 'vue';
import { useStore } from 'vuex';
import { useI18n } from 'vue-i18n'
import { BACKEND_SERVER, ICON_SERVER, SVG } from '~/constants';
import { ElMessage, TableColumnCtx } from 'element-plus';
import { Picture as IconPicture, TakeawayBox, Search } from '@element-plus/icons-vue'
import FileSaver from "file-saver";
import { utils, write } from 'xlsx';
import router from '~/router';

interface Material {
    ItemId: number
    MaterialName: string
    Price: number
    Quantity: number
    Cost: number
    IsBluePrint: boolean
    Err: boolean
    ErrMessage: string
    length: number
}

interface Item {
    ItemId: number
    ItemName: string
    Matertials: Material[]
    Quantity: number
    LpCost: number
    IskCost: number
    Cost: number
    MaterialCost: number
    Price: number
    Income: number
    Profit: number
    Volume: number
    SaleIndex: number
    UnitProfit: number
    Error: boolean
    ErrMessage: string
}

const store = useStore();
const { t, locale } = useI18n()

let tableData = reactive<Item[]>([])
const search = ref('')
const currentPage = ref(1)
const pageSize = ref(25)
const loading = ref(false)
let selectedItems: Item[] = []

interface SpanMethodProps {
    row: Material
    column: TableColumnCtx<Material>
    rowIndex: number
    columnIndex: number
}

const form = computed(() => store.state.form)
const calculator = computed(() => store.state.tableForm.calculator === "true");
const tableLayout = computed(() => store.state.tableForm.layout);
const highlight = computed(() => store.state.tableForm.highlight === "true");
const regionId = computed(() => store.state.regionId)
const corporationId = computed(() => store.state.corporationId)
const filterTableData = computed(() =>
    tableData.filter(
        (data) =>
            !search.value ||
            data.ItemName.toLowerCase().includes(search.value.toLowerCase())
    ).slice((currentPage.value - 1) * pageSize.value, currentPage.value * pageSize.value)
)

const pageWidth = ref(window.innerWidth);
const operationWidth = computed(() => {
    return pageWidth.value * 0.12;
});
window.addEventListener('resize', () => {
    pageWidth.value = window.innerWidth;
});

const sortFunctions = {
    LpCost: {
        ascending: (a: { LpCost: number; }, b: { LpCost: number; }) => a.LpCost - b.LpCost,
        descending: (a: { LpCost: number; }, b: { LpCost: number; }) => b.LpCost - a.LpCost,
    },
    IskCost: {
        ascending: (a: { IskCost: number; }, b: { IskCost: number; }) => a.IskCost - b.IskCost,
        descending: (a: { IskCost: number; }, b: { IskCost: number; }) => b.IskCost - a.IskCost,
    },
    MaterialCost: {
        ascending: (a: { MaterialCost: number; }, b: { MaterialCost: number; }) => a.MaterialCost - b.MaterialCost,
        descending: (a: { MaterialCost: number; }, b: { MaterialCost: number; }) => b.MaterialCost - a.MaterialCost,
    },
    Price: {
        ascending: (a: { Price: number; }, b: { Price: number; }) => a.Price - b.Price,
        descending: (a: { Price: number; }, b: { Price: number; }) => b.Price - a.Price,
    },
    Income: {
        ascending: (a: { Income: number; }, b: { Income: number; }) => a.Income - b.Income,
        descending: (a: { Income: number; }, b: { Income: number; }) => b.Income - a.Income,
    },
    Profit: {
        ascending: (a: { Profit: number; }, b: { Profit: number; }) => a.Profit - b.Profit,
        descending: (a: { Profit: number; }, b: { Profit: number; }) => b.Profit - a.Profit,
    },
    Volume: {
        ascending: (a: { Volume: number; }, b: { Volume: number; }) => a.Volume - b.Volume,
        descending: (a: { Volume: number; }, b: { Volume: number; }) => b.Volume - a.Volume,
    },
    SaleIndex: {
        ascending: (a: { SaleIndex: number; }, b: { SaleIndex: number; }) => a.SaleIndex - b.SaleIndex,
        descending: (a: { SaleIndex: number; }, b: { SaleIndex: number; }) => b.SaleIndex - a.SaleIndex,
    },
    UnitProfit: {
        ascending: (a: { UnitProfit: number; }, b: { UnitProfit: number; }) => a.UnitProfit - b.UnitProfit,
        descending: (a: { UnitProfit: number; }, b: { UnitProfit: number; }) => b.UnitProfit - a.UnitProfit,
    },

};

watch([form, regionId, corporationId, locale], () => {
    if (corporationId.value === undefined) {
        tableData.splice(0, tableData.length);
        return;
    }
    if (corporationId.value === 0) {
        return
    }
    loadTableData();
});

function loadTableData() {
    loading.value = true
    axios.get(`${BACKEND_SERVER}offer`, {
        params: {
            regionId: regionId.value,
            scope: form.value.scope,
            corporationId: corporationId.value,
            lang: locale.value,
            days: form.value.days,
            tax: form.value.tax,
            productPrice: form.value.productPrice,
            materialPrice: form.value.materialPrice,
        },
    })
        .then(response => {
            const data = response.data;
            for (let i = 0; i < data.length; ++i) {
                var matertials = data[i].Matertials;
                if (matertials == null) {
                    continue;
                }
                let count = 0;
                for (let j = 0; j + count < matertials.length;) {
                    if (count == 0) {
                        matertials[j].length = 1;
                        ++count;
                        continue;
                    }
                    if (matertials[j].IsBluePrint == matertials[j + count].IsBluePrint) {
                        matertials[j].length += 1;
                        matertials[j + count].length = 0;
                        ++count;
                    } else {
                        j += count;
                        count = 0;
                    }
                }
                tableData.splice(0, tableData.length, ...response.data);
            }
        })
        .catch(error => {
            console.error(error);
        })
        .finally(() => {
            loading.value = false
        });
}

const sortChange = (column: { prop: keyof typeof sortFunctions; order: 'ascending' | 'descending'; }) => {
    const sortFunction = sortFunctions[column.prop];
    if (sortFunction) {
        const orderFunction = column.order === 'ascending' ? sortFunction.ascending : sortFunction.descending;
        if (sortFunction) {
            const orderFunction = column.order === 'ascending' ? sortFunction.ascending : sortFunction.descending;
            const sortedData = [...tableData].sort(orderFunction).slice();
            tableData.splice(0, tableData.length, ...sortedData);
        }
    }
};

function getIcon(id: string): string {
    return `${ICON_SERVER}Type/${id}_32.png`;
}

function tableStyle() {
    return {
        padding: "0",
    }
}

function handleSizeChange(val: number) {
    currentPage.value = 1;
    pageSize.value = val;
}

function handleCurrentChange(val: number) {
    currentPage.value = val;
}

function numberFormat(row: Record<string, any>, column: TableColumnCtx<Record<string, any>>, cellValue: number): string {
    let formattedValue = Math.round(cellValue).toString();
    if (!formattedValue.includes(".")) formattedValue += ".";
    return formattedValue
        .replace(/(\d)(?=(\d{3})+\.)/g, ($0, $1) => {
            return $1 + ",";
        })
        .replace(/\.$/, "");
}

function objectSpanMethod({
    row,
    column,
    rowIndex,
    columnIndex,
}: SpanMethodProps) {
    if (columnIndex === 0) {
        if (row.length != 0) {
            return {
                rowspan: row.length,
                colspan: 1,
            };
        } else {
            return {
                rowspan: 0,
                colspan: 0,
            };
        }
    }
}

function handelRowDetail({ row, rowIndex }: { row: Item; rowIndex: number }): string {
    let value = '';
    if (row.Matertials === null) {
        value += "row-expand-cover ";
    }
    if (row.Error === true) {
        value += "warning-row";
    }
    return value.trim();
}


function exportExcel(corporationName: string) {
    if (tableData.length === 0) {
        ElMessage({
            message: 'No Data to export.',
            type: 'warning',
            duration: 1500,
        })
        return;
    }

    const pages = pageSize.value;
    pageSize.value = 500;
    currentPage.value = 1;
    let date = new Date();
    let name = corporationName + "-" + dateFormat("YYYYmmdd_HHMMSS", date) + "-" + form.value.tax + "%" + t('config.form.tax');
    nextTick(() => {
        let wb = utils.table_to_book(document.getElementById("table"));
        let wbout = write(wb, {
            bookType: "xlsx",
            bookSST: true,
            type: "array",
        });
        try {
            FileSaver.saveAs(
                new Blob([wbout], { type: "application/octet-stream" }),
                name + ".xlsx"
            );
        } catch (e) {
            console.error(e, wbout);
        }
        pageSize.value = pages;
    });
}

function dateFormat(fmt: string, date: Date): string {
    let ret: RegExpExecArray | null;
    const opt: { [key: string]: string } = {
        "Y+": date.getFullYear().toString(),
        "m+": (date.getMonth() + 1).toString(),
        "d+": date.getDate().toString(),
        "H+": date.getHours().toString(),
        "M+": date.getMinutes().toString(),
        "S+": date.getSeconds().toString(),
    };
    for (let k in opt) {
        ret = new RegExp("(" + k + ")").exec(fmt);
        if (ret) {
            fmt = fmt.replace(
                ret[1],
                ret[1].length === 1 ? opt[k] : opt[k].padStart(ret[1].length, "0")
            );
        }
    }
    return fmt;
}

function errorMessage(errMessage: string) {
    var lan = localStorage.lang;
    var message = errMessage;
    if (lan != 'en') {
        message = message
            .replace(/failed to get buy price for <b>(.*?)<\/b>/g, t('table.err.productBuy'))
            .replace(/failed to get sell price for <b>(.*?)<\/b>/g, t('table.err.productSell'))
            .replace(/failed to get buy price for production material <b>(.*?)<\/b>/g, t('table.err.materialBuy'))
            .replace(/failed to get sell price for production material <b>(.*?)<\/b>/g, t('table.err.materialSell'))
            .replace(/failed to get buy price for requirement <b>(.*?)<\/b>/g, t('table.err.requirementBuy'))
            .replace(/failed to get sell price for requirement <b>(.*?)<\/b>/g, t('table.err.requirementSell'))
            .replace(/no buy order found in the market/g, t('table.err.buyOrder'))
            .replace(/no sell order found in the market/g, t('table.err.sellOrder'))
            .replace(/no order found in the market/g, t('table.err.order'));
    }
    ElMessage({
        dangerouslyUseHTMLString: true,
        message: message,
        type: "warning",
        showClose: true,
        duration: 6000,
    });
}

function copyAllMaterials(list: any[]) {
    let value = list.map(item => `${item.Quantity} ${item.MaterialName}`).join("\n");
    navigator.clipboard.writeText(value).then(
        () => {
            ElMessage({
                message: t('table.material.copySuccess'),
                type: "success",
                duration: 2000,
            });
        },
        (error) => {
            ElMessage({
                message: t('table.material.copyFail'),
                type: "error",
                duration: 2000,
            });
            console.error(error);
        }
    );
}

function copyMaterial(row: Material) {
    var value = row.Quantity + " " + row.MaterialName + "\n";
    navigator.clipboard.writeText(value).then(
        () => {
            ElMessage({
                message: t('table.material.copySuccess'),
                type: "success",
                duration: 2000,
            });
        },
        (error) => {
            ElMessage({
                message: t('table.material.copyFail'),
                type: "error",
                duration: 2000,
            });
            console.error(error);
        }
    );
}

function formatNumber(value: number) {
    let result = value.toString();
    if (result.indexOf('.') !== -1) {
        let decimalLength = result.split('.')[1].length;
        if (decimalLength > 2) {
            return parseFloat(result).toFixed(2);
        }
    }
    return value;
}

function openItem(row: Item) {
    axios.get(`${BACKEND_SERVER}itemdetail`, {
        params: {
            itemId: row.ItemId,
            lang: locale.value
        },
        timeout: 5000,
    })
        .then((response) => {
            let ItemName = response.data.ItemName;
            let Volume = response.data.Volume;
            if (row.Quantity > 1) {
                ItemName = row.Quantity + "\u2009×\u2009" + ItemName
                Volume = Volume + "\u2009m³\u2009/\u2009" + formatNumber(Volume * row.Quantity)
            }
            Volume = Volume + "\u2009m³";
            let itemDetail = {
                ItemId: response.data.ItemId,
                ItemName: ItemName,
                Description: response.data.Description,
                Volume: Volume
            };
            store.commit('setItemVisible', true);
            store.commit('setItemDetail', itemDetail);
        })
        .catch(error => {
            let params = new URLSearchParams(error.config.params).toString()
            let fullUrl = error.config.url + (params ? `?${params}` : '');
            ElMessage({
                dangerouslyUseHTMLString: true,
                message: `${error.response.data}: ${fullUrl}`,
                type: "error",
                showClose: true,
                duration: 6000,
            });
        });
}

function handleSelectionChange(selection: Item[]) {
    selectedItems = selection
}

function openCalculator() {
    if (selectedItems.length == 0) {
        ElMessage({
            message: t('calculator.empty'),
            type: "warning",
            showClose: true,
            duration: 2000,
        });
        return
    }
    store.commit('setCalculatorVisible', true);
    store.commit('setSelectedItems', selectedItems);
}

function openOrder(scope: any) {
    const queryParams = {
        material_price: form.value.materialPrice,
        weighted_price: form.value.scope,
        tax: form.value.tax
    };

    const routeData = router.resolve({
        name: 'Order',
        params: {
            regionId: regionId.value,
            corporationId: corporationId.value,
            offerId: scope.row.OfferId,
        },
        query: queryParams
    });

    window.open(routeData.href, '_blank');
}

defineExpose({
    exportExcel
})

</script>
<style>
.el-table tr {
    background-color: transparent !important;
}

.el-table {
    --el-table-header-bg-color: transparent !important;
    --el-table-bg-color: transparent !important;
    caret-color: transparent !important;
}

.row-expand-cover .el-table__expand-column .cell {
    display: none !important;
}

.grab-cursor:hover {
    cursor: pointer;
}
</style>