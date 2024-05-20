<template>
    <div class="History" v-if="renderStock">
        <highcharts class="stock" :constructor-type="'stockChart'" :options="chartOptions"></highcharts>
    </div>
</template>

<script lang="ts" setup>
import axios from 'axios';
import { onBeforeMount, onMounted, reactive, ref, watch } from 'vue';
import { BACKEND_SERVER } from '~/constants';
import { useI18n } from 'vue-i18n'
import { useDark } from '@vueuse/core';

const { t, locale } = useI18n()
const isDark = useDark();
const props = defineProps<{
    offerId: string;
    regionId: string;
}>();

let renderStock = ref(false)
const history = reactive({
    average: [] as [number, number][],
    volume: [] as [number, number][],
    minAndmax: [] as [number, number, number][],
    average5d: [] as [number, number][],
    average20d: [] as [number, number][],
    minAndmax5d: [] as [number, number, number][],
    borderWidth: document.documentElement.clientWidth * 0.0388,
});


const chartColors = reactive({
    background: isDark.value ? "#202124" : "#f5f5f5",
    font: isDark.value ? "#D0D3D4" : "#000000",
    buttonHover: isDark.value ? "#0D4579" : "#f5f5f5",
    buttonSelect: isDark.value ? "#0D4579" : "#f5f5f5",
    menuBackground: isDark.value ? "#202124" : "#f5f5f5",
    menuItemHover: isDark.value ? "#0D4579" : "#f5f5f5",
    menuItemText: isDark.value ? "#D0D3D4" : "#000000",
    navigatorColor: isDark.value ? "#3498DB" : "#1561ac",
    xAxis: isDark.value ? "#505053" : "#ccd6eb",
    yAxis: isDark.value ? "#505053" : "#ccd6eb",
    tooltipBackground: isDark.value ? "rgba(0,0,0,0.8)" : "rgba(255,255,255,0.8)",
    tooltipText: isDark.value ? "#D0D3D4" : "#000000",
});

onMounted(async () => {
    getHistory()
});

async function getHistory() {
    try {
        const response = await axios.get(BACKEND_SERVER + "history", {
            params: {
                regionId: props.regionId,
                offerId: props.offerId,
            },
        });

        response.data.forEach((item: any) => {
            history.average.push([Date.parse(item.Date), item.Average]);
            history.volume.push([Date.parse(item.Date), item.Volume]);
            history.minAndmax.push([Date.parse(item.Date), item.Lowest, item.Highest]);
            history.average5d.push([Date.parse(item.Date), item.Average5d]);
            history.average20d.push([Date.parse(item.Date), item.Average20d]);
            history.minAndmax5d.push([Date.parse(item.Date), item.Lowest5d, item.Highest5d]);
        });
        renderStock.value = true

    } catch (error) {
        console.error('Error fetching history data:', error);
    }
}

let chartOptions = reactive<any>({
    chart: {
        backgroundColor: chartColors.background,
        alignTicks: false,
        marginRight: history.borderWidth,
    },
    credits: {
        enabled: false
    },
    scrollbar: {
        enabled: false,
    },
    legend: {
        enabled: true,
        align: 'center',
        verticalAlign: 'top',
        backgroundColor: chartColors.background,
        layout: 'horizontal',
        itemStyle: {
            color: chartColors.font
        },
    },
    navigator: {
        series: {
            data: history.average,
            color: chartColors.navigatorColor,
        },
        xAxis: {
            labels: {
                style: {
                    color: chartColors.font,
                    fontWeight: 'bold',
                    textOutline: 'none'
                }
            },
            gridLineWidth: 0,
            minorGridLineWidth: 0,
            dateTimeLabelFormats: {
                millisecond: '%Y.%m.%d',
                second: '%Y.%m.%d',
                minute: '%Y.%m.%d',
                hour: '%Y.%m.%d',
                day: '%Y.%m.%d',
                week: '%Y.%m.%d',
                month: '%Y.%m.%d',
                year: '%Y.%m.%d'
            },
        },
        yAxis: {
            gridLineWidth: 0,
            minorGridLineWidth: 0,
        },
    },
    navigation: {
        buttonOptions: {
            theme: {
                states: {
                    hover: {
                        fill: chartColors.buttonHover
                    },
                    select: {
                        stroke: chartColors.buttonSelect,
                        fill: chartColors.buttonSelect
                    }
                }
            }
        },
        menuItemHoverStyle: {
            background: chartColors.menuItemHover,
            color: chartColors.menuItemText,
        },
        menuItemStyle: {
            color: chartColors.menuItemText,
        },
        menuStyle: {
            background: chartColors.menuBackground,
        },
    },
    exporting: {
        enabled: true,
        buttons: {
            contextButton: {
                symbolY: 10,
                height: 20,
                menuItems: ['viewFullscreen', "printChart"],
                symbolStroke: chartColors.font,
                theme: {
                    fill: 'none',
                },
            },
        },
    },
    rangeSelector: {
        inputDateFormat: '%Y.%m.%d',
        inputEditDateFormat: '%Y.%m.%d',
        allButtonsEnabled: true,
        selected: 0,
        inputEnabled: true,
        inputStyle: {
            color: chartColors.font
        },
        labelStyle: {
            color: chartColors.font,
            fontWeight: 'bold'
        },
        buttons: [{
            type: 'month',
            count: 1,
            text: t('order.stock.rangeSelector.month'),
        }, {
            type: 'month',
            count: 3,
            text: t('order.stock.rangeSelector.threeMonths'),
        }, {
            type: 'month',
            count: 6,
            text: t('order.stock.rangeSelector.halfYear'),
        }, {
            type: 'ytd',
            text: t('order.stock.rangeSelector.yearToDay'),
        }, {
            type: 'year',
            count: 1,
            text: t('order.stock.rangeSelector.year'),
        }, {
            type: 'all',
            text: t('order.stock.rangeSelector.all'),
        }],
        buttonTheme: {
            fill: "none",
            stroke: "none",
            "stroke-width": 0,
            r: 8,
            style: {
                color: chartColors.font,
                fontWeight: "bold",
            },
            states: {
                hover: {
                    fill: chartColors.buttonHover,
                    style: {
                        color: chartColors.font,
                    },
                },
                select: {
                    fill: chartColors.buttonSelect,
                    style: {
                        color: chartColors.font,
                    },
                },
            },
        },
    },
    xAxis: {
        gridLineWidth: 0,
        gridLineColor: chartColors.xAxis,
        lineColor: chartColors.xAxis,
        labels: {
            style: {
                color: chartColors.font,
            },
        },
        crosshair: {
            dashStyle: 'dot',
        },
        dateTimeLabelFormats: {
            millisecond: '%Y.%m.%d',
            second: '%Y.%m.%d',
            minute: '%Y.%m.%d',
            hour: '%Y.%m.%d',
            day: '%Y.%m.%d',
            week: '%Y.%m.%d',
            month: '%Y.%m.%d',
            year: '%Y.%m.%d'
        },
    },
    yAxis: [
        {
            gridLineWidth: 0.3,
            gridLineColor: chartColors.yAxis,
            minorGridLineWidth: 0,
            startOnTick: false,
            endOnTick: false,
            labels: {
                align: "right",
                x: -3,
                style: {
                    color: chartColors.font,
                },
            },
            title: {
                text: t('order.stock.price'),
                style: {
                    color: chartColors.font,
                },
            },
            height: "60%",
            lineWidth: 2,
            lineColor: chartColors.yAxis,
            resize: {
                enabled: true,
            },
        },
        {
            gridLineWidth: 0.3,
            gridLineColor: chartColors.yAxis,
            minorGridLineWidth: 0,
            labels: {
                align: "right",
                x: -3,
                style: {
                    color: chartColors.font,
                },
            },
            title: {
                text: t('order.stock.volume'),
                style: {
                    color: chartColors.font,
                },
            },
            top: "65%",
            height: "35%",
            offset: 0,
            lineWidth: 2,
            lineColor: chartColors.yAxis,
        },
    ],
    tooltip: {
        split: false,
        xDateFormat: "%Y.%m.%d",
        backgroundColor: chartColors.tooltipBackground,
        shared: true,
        valueDecimals: 0,
        style: {
            fontSize: 13,
            color: chartColors.tooltipText,
        },
        headerFormat: '<span style="font-size: 13px">{point.key}</span><br/>',
    },
    series: [
        {
            name: t('order.stock.average'),
            data: history.average,
            yAxis: 0,
            color: "#D68910",
            lineWidth: 0,
            zIndex: 9,
            marker: {
                enabled: true,
                radius: 2.65,
            },
        },
        {
            name: t('order.stock.minAndmax'),
            data: history.minAndmax,
            type: "columnrange",
            yAxis: 0,
            zIndex: 8,
            color: "#A6ACAF",
            opacity: 0.6,
            pointWidth: 3,
        },
        {
            name: t('order.stock.average5d'),
            data: history.average5d,
            zIndex: 7,
            yAxis: 0,
            color: "#1D8348",
            lineWidth: 1.5,
            marker: {
                radius: 2.65,
            },
        },
        {
            name: t('order.stock.average20d'),
            data: history.average20d,
            zIndex: 7,
            yAxis: 0,
            color: "#C0392B",
            lineWidth: 1.5,
            marker: {
                radius: 1.65,
            },
        },
        {
            name: t('order.stock.minAndmax5d'),
            data: history.minAndmax5d,
            type: "arearange",
            zIndex: 6,
            yAxis: 0,
            color: "#85C1E9",
            opacity: 0.2,
        },
        {
            name: t('order.stock.volume'),
            type: "column",
            data: history.volume,
            yAxis: 1,
            color: "#0F5C70",
        },
    ],
});
</script>
<style scoped>
.stock {
    width: 100%;
    height: 100%;
}

input.highcharts-range-selector:focus {
    background-color: #0D4579;
}
</style>