<template>
    <div style="display: flex; justify-content: center; align-items: center;margin-top: 4%;">
        <el-image style="height: 48px; vertical-align: middle; margin-top: -22px;" :src="getIcon()" fit="contain" lazy>
            <div slot="error" class="image-slot">
                <i class="el-icon-picture-outline"></i>
            </div>
        </el-image>
        &nbsp;
        &nbsp;
        <div>
            <div style="display: flex; align-items: center; justify-content: center;">
                <h3 :style="{ color: titleColor, fontSize: '14px', textAlign: 'center', marginTop: '-5px' }">
                    {{ props.itemName }}
                </h3>
                &thinsp;
                <h3 :style="{ color: titleColor, fontSize: '14px', textAlign: 'center', marginTop: '-5px' }">
                    ({{ regionName }})
                </h3>
            </div>
            <h4 :style="{ color: titleColor, fontSize: '14px', textAlign: 'center', marginTop: '-5px' }">
                {{ props.corporationName }}
            </h4>
        </div>
    </div>
</template>
<script lang="ts" setup>
import { computed, onMounted, ref, watch } from 'vue';
import { ICON_SERVER } from '~/constants';
import { useI18n } from 'vue-i18n';
import { useDark } from '@vueuse/core'

const { t,locale } = useI18n();
const isDark = useDark()
const props = defineProps<{
    itemId: number,
    itemName: string,
    corporationName: string,
    regionId: string,
}>();
let regionName = ref('')
const titleColor = computed(() => {
  return isDark.value ? '#CDD0D1' : '#606266';
});

onMounted(() => {
    getRegionName()
})

watch(locale, () => {
    getRegionName()
})

function getIcon(): string {
    return `${ICON_SERVER}Type/${props.itemId}_64.png`;
}

function getRegionName() {
    switch (props.regionId) {
        case "10000002":
            regionName.value = t('header.region.TheForge');
            break
        case '10000043':
            regionName.value = t('header.region.Domain');
            break
        case '10000030':
            regionName.value = t('header.region.Heimatar');
            break
        case '10000032':
            regionName.value = t('header.region.SinqLaison');
            break
    }
}

</script>