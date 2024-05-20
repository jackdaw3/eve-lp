<template>
    <div class="dialog-container" style="display: flex; justify-content: space-between;">
        <div style="flex-grow: 1; display: flex; justify-content: flex-start;">
            <el-button :icon="Setting" circle size="large" @click="visible = true"></el-button>&nbsp;&nbsp;
            <el-dialog :title="t('config.title')" v-model="visible" width="38%" @close="closeDialog">
                <el-tabs model-value="first" type="card">
                    <el-tab-pane :label="t('config.data')" name="first">
                        <el-form label-width="20%" style="margin-top: 2.5%" v-model="form">
                            <el-form-item :label="t('config.form.materialPrice')">
                                <el-select :placeholder="t('config.form.materialPlaceholder')"
                                    v-model="form.materialPrice" style="width: 90%" :suffix-icon="SoldOut">
                                    <el-option :label="t('config.form.buyPrice')" value="buy"></el-option>
                                    <el-option :label="t('config.form.sellPrice')" value="sell"></el-option>
                                </el-select>
                            </el-form-item>
                            <el-form-item :label="t('config.form.productPrice')">
                                <el-select :placeholder="t('config.form.productPlaceholder')"
                                    v-model="form.productPrice" style="width: 90%" :suffix-icon="Sell">
                                    <el-option :label="t('config.form.buyPrice')" value="buy"></el-option>
                                    <el-option :label="t('config.form.sellPrice')" value="sell"></el-option>
                                </el-select>
                            </el-form-item>
                            <el-form-item :label="t('config.form.scope')">
                                <el-select :placeholder="t('config.form.scopePlaceholder')" v-model="form.scope"
                                    style="width: 90%" :suffix-icon="Coin">
                                    <el-option label="1%" value="0.01"></el-option>
                                    <el-option label="5%" value="0.05"></el-option>
                                    <el-option label="10%" value="0.1"></el-option>
                                    <el-option label="25%" value="0.25"></el-option>
                                </el-select>
                            </el-form-item>
                            <el-form-item :label="t('config.form.days')">
                                <el-select :placeholder="t('config.form.daysPlaceholder')" v-model="form.days"
                                    style="width: 90%" :suffix-icon="Timer">
                                    <el-option :label="t('config.form.week')" value="7"></el-option>
                                    <el-option :label="t('config.form.month')" value="30"></el-option>
                                    <el-option :label="t('config.form.season')" value="90"></el-option>
                                </el-select>
                            </el-form-item>
                            <el-form-item :label="t('config.form.tax')">
                                <el-slider v-model="form.tax" :max="20" show-input :format-tooltip="taxFormat"
                                    style="width: 90%"></el-slider>
                            </el-form-item>
                        </el-form>
                    </el-tab-pane>
                    <el-tab-pane :label="t('config.tableForm.title')" name="second">
                        <el-form label-width="20%" style="margin-top: 2.5%" v-model="form">
                            <el-form-item :label="t('config.tableForm.calculator')">
                                <el-select :placeholder="t('config.tableForm.calculatorPlaceholder')" v-model="tableForm.calculator"
                                    style="width: 90%" :suffix-icon="TakeawayBox">
                                    <el-option :label="t('config.tableForm.calculatorEnable')" value="true"></el-option>
                                    <el-option :label="t('config.tableForm.calculatorDisable')" value="false"></el-option>
                                </el-select>
                            </el-form-item>
                            <el-form-item :label="t('config.tableForm.layout')">
                                <el-select :placeholder="t('config.tableForm.calculatorPlaceholder')" v-model="tableForm.layout"
                                    style="width: 90%" :suffix-icon="List">
                                    <el-option :label="t('config.tableForm.layoutAuto')" value="auto"></el-option>
                                    <el-option :label="t('config.tableForm.layoutFixed')" value="fixed"></el-option>
                                </el-select>
                            </el-form-item>
                            <el-form-item :label="t('config.tableForm.highlight')">
                                <el-select :placeholder="t('config.tableForm.highlightPlaceholder')" v-model="tableForm.highlight"
                                    style="width: 90%" :suffix-icon="Opportunity">
                                    <el-option :label="t('config.tableForm.highlightEnable')" value="true"></el-option>
                                    <el-option :label="t('config.tableForm.highlightDisable')" value="false"></el-option>
                                </el-select>
                            </el-form-item>
                        </el-form>
                    </el-tab-pane>
                    <el-tab-pane :label="t('config.quickBar')" name="third">
                        <el-transfer
                            style="display: flex; align-items: center; justify-content: center; margin-top: 2.5%"
                            v-model="quickBarSelected" filterable
                            :titles="[t('config.sourceList'), t('config.targetList')]" :data="qucikBarAll">
                        </el-transfer>
                    </el-tab-pane>
                    <el-tab-pane :label="t('config.desc.title')" name="fourth">
                        <el-collapse accordion>
                            <el-collapse-item :title="t('config.desc.dataDesc')" name="1">
                                <el-descriptions class="margin-top" :column="1" border size="small">
                                    <el-descriptions-item>
                                        <template #label>
                                            <div class="cell-item">
                                                {{ t('config.form.materialPrice') }}
                                            </div>
                                        </template>
                                        {{ t('config.desc.materialContent') }}
                                    </el-descriptions-item>
                                    <el-descriptions-item>
                                        <template #label>
                                            <div class="cell-item">
                                                {{ t('config.form.productPrice') }}
                                            </div>
                                        </template>
                                        {{ t('config.desc.productPriceContent') }}
                                    </el-descriptions-item>
                                    <el-descriptions-item>
                                        <template #label>
                                            <div class="cell-item">
                                                {{ t('config.form.scope') }}
                                            </div>
                                        </template>
                                        {{ t('config.desc.scopeContent') }}
                                    </el-descriptions-item>
                                    <el-descriptions-item>
                                        <template #label>
                                            <div class="cell-item">
                                                {{ t('config.form.days') }}
                                            </div>
                                        </template>
                                        {{ t('config.desc.daysContent') }}
                                    </el-descriptions-item>
                                    <el-descriptions-item>
                                        <template #label>
                                            <div class="cell-item">
                                                {{ t('config.form.tax') }}
                                            </div>
                                        </template>
                                        {{ t('config.desc.taxContent') }}
                                    </el-descriptions-item>
                                </el-descriptions>
                            </el-collapse-item>
                            <el-collapse-item :title="t('config.desc.tableDesc')" name="2">
                                <el-descriptions class="margin-top" :column="1" border size="small">
                                    <el-descriptions-item>
                                        <template #label>
                                            <div class="cell-item" style="white-space: nowrap;">{{
                                                t('table.materialCost') }}</div>
                                        </template>
                                        {{ t('config.desc.costContent') }}
                                    </el-descriptions-item>
                                    <el-descriptions-item>
                                        <template #label>
                                            <div class="cell-item" style="white-space: nowrap;">{{ t('table.income') }}
                                            </div>
                                        </template>
                                        {{ t('config.desc.incomeContent') }}
                                    </el-descriptions-item>
                                    <el-descriptions-item>
                                        <template #label>
                                            <div class="cell-item" style="white-space: nowrap;">{{ t('table.volume') }}
                                            </div>
                                        </template>
                                        {{ t('config.desc.volumeContent') }}
                                    </el-descriptions-item>
                                    <el-descriptions-item>
                                        <template #label>
                                            <div class="cell-item" style="white-space: nowrap;">{{ t('table.saleIndex')
                                                }}</div>
                                        </template>
                                        {{ t('config.desc.saleIndexContent') }}
                                    </el-descriptions-item>
                                    <el-descriptions-item>
                                        <template #label>
                                            <div class="cell-item" style="white-space: nowrap;">{{ t('table.unitProfit')
                                                }}</div>
                                        </template>
                                        {{ t('config.desc.unitProfitContent') }}
                                    </el-descriptions-item>
                                    <el-descriptions-item>
                                        <template #label>
                                            <div class="cell-item" style="white-space: nowrap;">{{
                                                t('config.desc.blueprint') }}</div>
                                        </template>
                                        {{ t('config.desc.blueprintContent') }}
                                    </el-descriptions-item>
                                </el-descriptions>
                            </el-collapse-item>
                            <el-collapse-item :title="t('config.desc.claim')" name="3">
                                <div style="word-break: keep-all;">
                                    {{ t('config.desc.claimContent') }}
                                </div>
                            </el-collapse-item>
                        </el-collapse>
                    </el-tab-pane>
                </el-tabs>
                <template #footer>
                    <el-button style="max-width: 30px;" size="default" @click="reset">{{ t('config.reset')
                        }}</el-button>
                    <el-button style="max-width: 30px;" size="default" @click="visible = false">{{ t('config.close')
                        }}</el-button>
                </template>
            </el-dialog>
            <el-cascader style="width: 30%" :options="options" v-model="corporationIds" size="large"
                :placeholder="t('config.corporation')" clearable filterable>
                <template #default="{ node, data }">
                    <span>
                        <img style="height: 22px; vertical-align: middle; margin-top: -3px;" :src="data.imageUrl"
                            loading="lazy">
                        &nbsp;
                        {{ node.label }}
                    </span>
                </template>
            </el-cascader>
        </div>
        <div style="flex-grow: 1; display: flex; justify-content: flex-end; margin-right: 15px;">
            <el-button :icon="Download" circle size="large" style="cursor: pointer" @click="downloadExcel" />
        </div>

    </div>
</template>

<script lang="ts" setup>
import { onMounted, reactive, ref, watch, defineEmits, toRaw } from 'vue'
import { useI18n } from 'vue-i18n'
import { Setting } from '@element-plus/icons-vue'
import { useStore } from 'vuex';
import { BACKEND_SERVER, ICON_SERVER } from '~/constants';
import axios from 'axios';
import { Download, SoldOut, Sell, Coin, Timer, TakeawayBox, List, Opportunity } from '@element-plus/icons-vue'
import { useDark } from '@vueuse/core'

const { t, locale } = useI18n()
const store = useStore();
const emit = defineEmits(['exportRequest'])

const visible = ref<boolean>(false)
const options = ref<any>([])
const corporationIds = ref()
const isDark = useDark()
interface Form {
    materialPrice: string;
    productPrice: string;
    scope: string;
    days: string;
    tax: number;
}
let form = reactive<Form>({
    materialPrice: 'sell',
    productPrice: 'buy',
    scope: "0.05",
    days: "7",
    tax: 0
});

interface TableForm {
    calculator: string;
    highlight: string;
    layout: string;
}
let tableForm = reactive<TableForm>({
    calculator: "false",
    highlight: "false",
    layout: "auto",
});

type QucikBarItem = {
    key: number;
    label: string;
};
const qucikBarAll = reactive<QucikBarItem[]>([]);
let quickBarSelected = ref([])

onMounted(() => {
    loadFactions();
    const storedQuickBarSelected = localStorage.getItem('quickBarSelected');
    if (storedQuickBarSelected) {
        quickBarSelected.value = JSON.parse(storedQuickBarSelected);
    }

    const storedForm = localStorage.getItem('form');
    if (storedForm) {
        form = reactive<Form>(JSON.parse(storedForm));
        store.commit('setForm', { ...form });
    }

    const storedTableForm = localStorage.getItem('tableForm');
    if (storedTableForm) {
        tableForm = reactive<TableForm>(JSON.parse(storedTableForm));
        store.commit('setTableForm', { ...tableForm });
    }
});

watch(corporationIds, (newCorporationIds) => {
    if (newCorporationIds) {
        store.commit('setCorporationId', parseInt(newCorporationIds[1]));
    } else {
        store.commit('setCorporationId', undefined);
    }
})

watch(locale, () => {
    loadFactions();
})

watch(isDark, () => {
    reloadQuickbar();
})

async function loadFactions() {
    try {
        const response = await axios.get(`${BACKEND_SERVER}faction`, {
            params: {
                lang: locale.value,
            },
        });
        const factions = response.data;
        loadCorporations(factions);
        processCorporationsData(factions);
    } catch (error) {
        console.error(error);
    }
}

function loadCorporations(list: any[]) {
    const factions: any[] = list.map(function (factionItem) {
        const faction = {
            value: factionItem.FactionId,
            label: factionItem.FactionName,
            imageUrl: getIcon(factionItem.FactionId),
            children: factionItem.Corporations ? factionItem.Corporations.map(function (corporationItem: { CorporationId: any; CorporationName: any; }) {
                return {
                    value: corporationItem.CorporationId,
                    label: corporationItem.CorporationName,
                    imageUrl: getIcon(corporationItem.CorporationId)
                };
            }) : []
        };
        return faction;
    });
    options.value = factions;
    reloadQuickbar();
}

function processCorporationsData(originalData: any[]) {
    qucikBarAll.splice(0, qucikBarAll.length);
    originalData.forEach(function (faction) {
        faction.Corporations.forEach(function (corporation: { CorporationId: number; CorporationName: string; }) {
            qucikBarAll.push({
                key: corporation.CorporationId,
                label: corporation.CorporationName
            });
        });
    });
}

function closeDialog() {
    function quickBarChange() {
        const storedQuickBarSelected = localStorage.getItem('quickBarSelected');
        if (!storedQuickBarSelected && quickBarSelected.value.length != 0) {
            return true
        }
        if (!storedQuickBarSelected) {
            return false
        }
        const oldQuickBarSelected = JSON.parse(storedQuickBarSelected);
        if (oldQuickBarSelected.length != quickBarSelected.value.length) {
            return true
        }
        for (let i = 0; i < oldQuickBarSelected.length; i++) {
            if (oldQuickBarSelected[i] != quickBarSelected.value[i]) {
                return true
            }
        }
        return false
    }

    function formchange() {
        const oldForm = store.state.form
        if (form.materialPrice != oldForm.materialPrice || form.productPrice != oldForm.productPrice
            || form.days != oldForm.days || form.tax != oldForm.tax || form.scope != oldForm.scope) {
            return true
        }
        return false
    }

    function tableFormChange() {
        const oldTableForm = store.state.tableForm
        if (tableForm.calculator != oldTableForm.calculator || tableForm.layout != oldTableForm.layout
            || tableForm.highlight != oldTableForm.highlight) {
            return true
        }
        return false
    }

    if (quickBarChange()) {
        localStorage.setItem('quickBarSelected', JSON.stringify(quickBarSelected.value));
        reloadQuickbar();
    }

    if (formchange()) {
        localStorage.setItem('form', JSON.stringify({ ...form }));
        store.commit('setForm', { ...form });
    }

    if (tableFormChange()) {
        localStorage.setItem('tableForm', JSON.stringify({ ...tableForm }));
        store.commit('setTableForm', { ...tableForm });
    }
}

function reloadQuickbar() {
    if (quickBarSelected.value === undefined) {
        return;
    }

    if (quickBarSelected.value.length > 0) {
        let collection: { value: number; label: string; imageUrl: string; children?: any[] } = {
            value: 1,
            label: t('config.quickBar'),
            imageUrl: isDark.value ? '/quickBar-dark.png' : '/quickBar-light.png',
        };
        let lists: any[] = [];
        for (let i = 0; i < quickBarSelected.value.length; ++i) {
            let temp: { value: string; label: string | undefined; imageUrl: string } = {
                value: quickBarSelected.value[i],
                label: getCorporationName(quickBarSelected.value[i]),
                imageUrl: getIcon(quickBarSelected.value[i])
            };
            lists.push(temp);
        }
        lists.sort((a, b) => a.value - b.value);
        collection.children = lists;
        if (options.value[0] && options.value[0].value === 1) {
            options.value.shift();
        }
        options.value.unshift(collection);
    } else if (quickBarSelected.value.length === 0 && options.value[0].value === 1) {
        options.value.shift();
    }
}

const getCorporationName = (itemId: string) => {
    for (let i = 0; i < options.value.length; ++i) {
        let faction = options.value[i];
        if (faction.children) {
            for (let j = 0; j < faction.children.length; ++j) {
                if (faction.children[j].value === itemId) {
                    return faction.children[j].label;
                }
            }
        }
    }
    return undefined;
}

function getIcon(id: string): string {
    return `${ICON_SERVER}Corporation/${id}_32.png`;
}

function taxFormat(e: number): string {
    return `${e}%`;
}

function downloadExcel() {
    emit('exportRequest', getCorporationName(corporationIds.value[1]))
}

function reset() {
    form.materialPrice = "sell"
    form.productPrice = "buy"
    form.scope = "0.05"
    form.days = "7"
    form.tax = 0
    store.commit('setForm', { ...form });
    localStorage.setItem('form', JSON.stringify({ ...form }));

    tableForm.calculator = "false"
    tableForm.layout = "auto"
    tableForm.highlight = "false"
    store.commit('setTableForm', { ...tableForm });
    localStorage.setItem('tableForm', JSON.stringify({ ...tableForm }));

    quickBarSelected.value = []
    localStorage.setItem('quickBarSelected', JSON.stringify(quickBarSelected.value));
}

</script>
<style>
.el-transfer-panel {
    width: 40% !important;
}

.el-transfer .el-transfer__button {
    width: 40%;
    padding: 0 15px;
}

.el-dropdown {
    font-size: 16px !important;
}

.cell-item {
    white-space: nowrap;
}
</style>