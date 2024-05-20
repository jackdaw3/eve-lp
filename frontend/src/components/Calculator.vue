<template>
    <div class="calculator">
        <el-dialog v-model="calculatorVisible" :title="t('calculator.title')" width="38%" draggable @close="resetAll">
            <el-form :model="selectedItems" label-position="left" label-width="auto"
                style="display: column; justify-content: center;">
                <el-form-item v-for="(item, index) in selectedItems" :key="item.ItemId" :label="item.ItemName"
                    class="selected-items">
                    <template #label>
                        <el-image :src="getIcon(item.ItemId)"
                            style="height: 22px; vertical-align: middle;margin-top: 6px"></el-image>&nbsp;
                        {{ getItemName(item) }}
                    </template>
                    <el-input-number v-model.number="inputValues[index]" style="margin-left:74%;" size="small"
                        :min="0"></el-input-number>
                </el-form-item>
                <br>
                <el-form-item>
                    <div style="display: flex;">
                        <el-button plain style="max-width: 30px;" @click="calculate"> {{ t('calculator.calculate')
                            }}</el-button>
                        <el-button plain style="max-width: 30px;" @click="resetAll"> {{ t('calculator.reset')
                            }}</el-button>
                    </div>

                </el-form-item>
            </el-form>

            <el-collapse v-model="materialListVisible">
                <el-collapse-item :title="t('calculator.materialList')" style="font-size: 14px;" name="1">
                    <el-table :data="materials" table-layout="auto" :cell-style="tableStyle"
                        :span-method="rowSpanMethod" :row-class-name="handelRowDetail"
                        :header-cell-style="{ padding: '0' }">
                        <el-table-column prop="IsBluePrint" align="center">
                            <template #default="scope">
                                <span v-if="scope.row.IsBluePrint === true">
                                    {{ t('table.material.bluePrintMaterial') }}
                                </span>
                                <span v-else>{{ t('table.material.lpStoreMaterial') }}</span>
                            </template>
                        </el-table-column>
                        <el-table-column prop="MaterialName" :label="t('table.material.name')">
                            <template #default="scope">
                                <img style="height: 22px; vertical-align: middle" :src="getIcon(scope.row.ItemId)"
                                    loading="lazy">
                                </img>
                                <span>{{ scope.row.MaterialName }}</span>
                            </template>
                        </el-table-column>
                        <el-table-column prop="Quantity" :label="t('table.material.quantity')"
                            :formatter="numberFormat" />
                        <el-table-column prop="Price" :label="t('table.material.price')" :formatter="numberFormat" />
                        <el-table-column prop="Cost" :label="t('table.material.cost')" :formatter="numberFormat" />
                        <el-table-column>
                            <template #header>
                                <el-button size="small" @click="copyAllMaterials(materials)" plain>{{
                                    t('table.material.copy') }}</el-button>
                            </template>
                            <template #default="scope">
                                <el-button size="small" @click="copyMaterial(scope)" plain>{{
                                    t('table.material.copy') }}</el-button>
                                <el-button v-if="scope.row.Error === true" size="small" plain
                                    @click="errorMessage(scope.row.ErrorMessage)" style="margin-left: 0px">{{
                                        t('table.material.error')
                                    }}</el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                </el-collapse-item>
            </el-collapse>

            <el-row style="margin-top: 18px;">
                <el-col :span="4">
                    <el-statistic :title="t('calculator.lpCost')" :value="statistic.lpCost" />
                </el-col>
                <el-col :span="5">
                    <el-statistic :title="t('calculator.iskCost')" :value="statistic.iskCost" />
                </el-col>
                <el-col :span="5">
                    <el-statistic :title="t('calculator.materialCost')" :value="statistic.materialCost" />
                </el-col>
                <el-col :span="5">
                    <el-statistic :title="t('calculator.profit')" :value="statistic.profit" />
                </el-col>
                <el-col :span="5">
                    <el-statistic :title="t('calculator.unitProfit')" :value="statistic.unitProfit" />
                </el-col>
            </el-row>

            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="calculatorVisible = false" style="max-width: 30px;" plain>
                        {{ t('calculator.close') }}
                    </el-button>
                </div>
            </template>

        </el-dialog>
    </div>
</template>

<script lang="ts" setup>
import { computed, reactive, ref, watch } from 'vue';
import { useStore } from 'vuex';
import { useI18n } from 'vue-i18n';
import { ICON_SERVER } from '~/constants';
import { ElMessage, TableColumnCtx } from 'element-plus';

const store = useStore();
const { t, locale } = useI18n();

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

interface SpanMethodProps {
    row: Material
    column: TableColumnCtx<Material>
    rowIndex: number
    columnIndex: number
}

interface Statistic {
    lpCost: number
    iskCost: number
    materialCost: number
    profit: number
    unitProfit: number
}

const calculatorVisible = computed({
    get: () => store.state.calculatorVisible,
    set: (value) => store.commit('setCalculatorVisible', value)
});
let selectedItems = computed(() => store.state.selectedItems);
let inputValues = ref(selectedItems.value.map(() => 0));

let statistic = reactive<Statistic>({
    lpCost: 0,
    iskCost: 0,
    materialCost: 0,
    profit: 0,
    unitProfit: 0
});
let materials = reactive<Material[]>([])
let materialListVisible = ref([])

watch(selectedItems, (newItems) => {
    inputValues.value = newItems.map(() => 1);
});

function getItemName(item: any): string {
    return item.Quantity + "\u2009×\u2009" + item.ItemName
}

function getIcon(id: string): string {
    return `${ICON_SERVER}Type/${id}_32.png`;
}



function processMaterials() {
    materials.splice(0, materials.length);
    type MaterialMap = {
        [key: number]: Material;
    };

    let materialsMap: MaterialMap = {};

    selectedItems.value.forEach((item: Item, index: number) => {
        if (!item.Matertials) {
            return;
        }

        const materialsWithResetLength = item.Matertials.map(material => ({
            ...material,
            length: 0,
            Quantity: material.Quantity * inputValues.value[index],
            Cost: material.Cost * inputValues.value[index]
        }));

        const materialsWithFilter = materialsWithResetLength.filter(material => material.Quantity !== 0);

        materialsWithFilter.forEach(material => {
            if (materialsMap[material.ItemId]) {
                materialsMap[material.ItemId].Quantity += material.Quantity;
                materialsMap[material.ItemId].Cost += material.Cost;
            } else {
                materialsMap[material.ItemId] = material;
            }
        });
    });


    materials.splice(0, materials.length, ...Object.values(materialsMap));
    materials.sort((a, b) => Number(a.IsBluePrint) - Number(b.IsBluePrint));
    let count = 0;
    for (let j = 0; j + count < materials.length;) {
        if (count == 0) {
            materials[j].length = 1;
            ++count;
            continue;
        }
        if (materials[j].IsBluePrint == materials[j + count].IsBluePrint) {
            materials[j].length += 1;
            materials[j + count].length = 0;
            ++count;
        } else {
            j += count;
            count = 0;
        }
    }
}

function resetStatistic() {
    statistic.lpCost = 0
    statistic.iskCost = 0
    statistic.materialCost = 0
    statistic.profit = 0
    statistic.unitProfit = 0
}

function resetAll() {
    for (let i = 0; i < inputValues.value.length; i++) {
        inputValues.value[i] = 1
    }
    materialListVisible.value = []
    materials.splice(0, materials.length);
    resetStatistic()
}

function calculate() {
    processMaterials()
    resetStatistic()
    for (let i = 0; i < selectedItems.value.length; i++) {
        statistic.lpCost += selectedItems.value[i].LpCost * inputValues.value[i]
        statistic.iskCost += selectedItems.value[i].IskCost * inputValues.value[i]
        statistic.materialCost += selectedItems.value[i].MaterialCost * inputValues.value[i]
        statistic.profit += selectedItems.value[i].Profit * inputValues.value[i]
    }
    statistic.unitProfit = statistic.profit / statistic.lpCost
}

function tableStyle() {
    return {
        padding: "0",
        "font-size": "14px",
    }
}

function rowSpanMethod({
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

function numberFormat(row: Record<string, any>, column: TableColumnCtx<Record<string, any>>, cellValue: number): string {
    let formattedValue = Math.round(cellValue).toString();
    if (!formattedValue.includes(".")) formattedValue += ".";
    return formattedValue
        .replace(/(\d)(?=(\d{3})+\.)/g, ($0, $1) => {
            return $1 + ",";
        })
        .replace(/\.$/, "");
}

function errorMessage(errMessage: string) {
    var lan = localStorage.lang;
    var message = errMessage;
    if (lan != 'en') {
        message = message
            .replace(/failed to get buy price for <b>(.*?)<\/b>/, t('table.err.produceBuy'))
            .replace(/failed to get sell price for <b>(.*?)<\/b>/, t('table.err.produceSell'))
            .replace(/failed to get buy price for production material <b>(.*?)<\/b>/, t('table.err.materialBuy'))
            .replace(/failed to get sell price for production material <b>(.*?)<\/b>/, t('table.err.materialSell'))
            .replace(/failed to get buy price for requirement <b>(.*?)<\/b>/, t('table.err.requirementBuy'))
            .replace(/failed to get sell price for requirement <b>(.*?)<\/b>/, t('table.err.requirementSell'))
            .replace('no buy order found in the market', t('table.err.buyOrder'))
            .replace('no sell order found in the market', t('table.err.sellOrder'))
            .replace('no order found in the market', t('table.err.order'));
    }

    ElMessage({
        dangerouslyUseHTMLString: true,
        message: message,
        type: 'warning',
        showClose: true,
        duration: 6000,
    });
}

function copyAllMaterials(list: any[]) {
    let value = list.map(function (item) { return `${item.Quantity} ${item.MaterialName}`; }).join('\n');
    navigator.clipboard.writeText(value).then(
        function () {
            ElMessage({
                message: t('table.material.copySuccess'),
                type: 'success',
                duration: 2000,
            });
        },
        function (error) {
            ElMessage({
                message: t('table.material.copyFail'),
                type: 'error',
                duration: 2000,
            });
            console.error(error);
        }
    );
}

function copyMaterial(row: any) {
    var value = row.Quantity + ' ' + row.MaterialName + '\n';
    navigator.clipboard.writeText(value).then(
        function () {
            ElMessage({
                message: t('table.material.copySuccess'),
                type: 'success',
                duration: 2000,
            });
        },
        function (error) {
            ElMessage({
                message: t('table.material.copyFail'),
                type: 'error',
                duration: 2000,
            });
            console.error(error);
        }
    );
}
</script>
<style>
.el-col {
    text-align: center;
}

.el-collapse-item__header {
    font-size: 14px !important;
}

.selected-items {
    margin-bottom: -6px;
}
</style>
