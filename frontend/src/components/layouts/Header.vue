<template>
  <div class="header">
    <el-menu :class="{ 'dark-mode': isDark }" class="menu" mode="horizontal" :ellipsis="false">
      <el-menu-item index="1" class="icon">
        <img style="width: 26px; vertical-align: middle; margin-left: 25px;margin-top: -2px;" src="/icon.png"
          class="logo" />&nbsp;&nbsp;
        <span style="font-family: 'Microsoft YaHei'; font-size: 20px; margin-right: 10px;">EVE-LP</span>
      </el-menu-item>

      <el-menu-item index="2" style="margin-left: -25px;">
        <el-dropdown>
          <span class="el-dropdown-link">
            {{ serverLabel }}
            <el-icon class="el-icon--right">
              <arrow-down />
            </el-icon>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="goToSerenity">{{ $t('header.server.Serenity') }}</el-dropdown-item>
              <el-dropdown-item @click="goToTranquility">{{ $t('header.server.Tranquility') }}</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </el-menu-item>
      <div class="flex-grow" />

      <el-menu-item index="3" style="margin-right: -25px;" v-if="regionVisible">
        <el-dropdown @command="regionChange">
          <span class="el-dropdown-link">
            {{ regionLabel }}
            <el-icon class="el-icon--right">
              <arrow-down />
            </el-icon>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item command="10000002">{{ $t('header.region.TheForge') }}</el-dropdown-item>
              <el-dropdown-item command="10000043">{{ $t('header.region.Domain') }}</el-dropdown-item>
              <el-dropdown-item command="10000032">{{ $t('header.region.SinqLaison') }}</el-dropdown-item>
              <el-dropdown-item command="10000030">{{ $t('header.region.Heimatar') }}</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </el-menu-item>

      <el-menu-item index="4" style="margin-right: -30px;">
        <el-dropdown @command="langChange">
          <span class="el-dropdown-link">
            {{ languageLabel }}
            <el-icon class="el-icon--right">
              <arrow-down />
            </el-icon>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item command="de">Deutsch</el-dropdown-item>
              <el-dropdown-item command="en">English</el-dropdown-item>
              <el-dropdown-item command="fr">Français</el-dropdown-item>
              <el-dropdown-item command="ja">日本語</el-dropdown-item>
              <el-dropdown-item command="ru">Pусский</el-dropdown-item>
              <el-dropdown-item command="zh">中文</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </el-menu-item>

      <el-menu-item h="full" @click="toggleDark()" style="margin-right: -20px;">
        <button class="border-none w-full bg-transparent cursor-pointer" style="height: var(--ep-menu-item-height)">
          <i inline-flex i="dark:ep-moon ep-sunny" />
        </button>
      </el-menu-item>

      <el-menu-item index="5">
        <img style="width: 21px; vertical-align: middle;" :src="iconSrc"
           @click="goToGithub" />
      </el-menu-item>

    </el-menu>
  </div>
</template>
<script lang="ts" setup>
import { ArrowDown } from '@element-plus/icons-vue'
import { ref, watch, computed, onBeforeMount } from 'vue'
import { useStore } from 'vuex';
import { useI18n } from 'vue-i18n'
import { useDark, useToggle } from '@vueuse/core'

type Language = 'de' | 'en' | 'fr' | 'ja' | 'ru' | 'zh';
type Region = '10000002' | '10000043' | '10000030' | '10000032';

const { t, locale } = useI18n()
const store = useStore();

const languageLabel = ref('')
const regionLabel = ref('')
const regionId = ref<Region>('10000002')
const serverLabel = ref('')
const isDark = useDark()
const toggleDark = useToggle(isDark)

defineProps<{
  regionVisible: boolean;
}>();

onBeforeMount(() => {
  const storedLanguage = localStorage.getItem('language')
  if (storedLanguage) {
    locale.value = storedLanguage as Language;
  }
  languageLabel.value = langLabel(locale.value as Language);
  regionLabel.value = regLabel(regionId.value);
  serverLabel.value = (t('header.server.Serenity'))
})

const iconSrc = computed(() => {
  return isDark.value ? '/github-dark.png' : '/github-light.png';
})

watch(locale, (newLocale) => {
  languageLabel.value = langLabel(newLocale as Language);
  localStorage.setItem('language', newLocale);
  regionLabel.value = regLabel(regionId.value);
  serverLabel.value = (t('header.server.Tranquility'))
})

watch(regionId, (newRegionId) => {
  regionLabel.value = regLabel(newRegionId);
})

function langLabel(lang: Language): string {
  switch (lang) {
    case 'de':
      return 'Deutsch';
    case 'en':
      return 'English';
    case 'fr':
      return 'Français';
    case 'ja':
      return '日本語';
    case 'ru':
      return 'Pусский';
    case 'zh':
      return '中文';
  }
}

function langChange(lang: Language) {
  locale.value = lang;
}

function regionChange(reg: Region) {
  regionId.value = reg;
  regionLabel.value = regLabel(reg);
  store.commit('setRegionId', reg);
}

function regLabel(reg: Region): string {
  switch (reg) {
    case '10000002':
      return t('header.region.TheForge');
    case '10000043':
      return t('header.region.Domain');
    case '10000030':
      return t('header.region.Heimatar');
    case '10000032':
      return t('header.region.SinqLaison');
  }
}

function goToTranquility() {
  window.location.href = 'https://eve-lp.com';
}

function goToSerenity() {
  window.location.href = 'https://eve-lp.com/serenity';
}

function goToGithub() {
  window.open("https://github.com/jackjaw/eve-lp", '_blank')
}

</script>

<style>
.el-menu--horizontal .el-menu-item:not(.is-disabled):hover,
.el-menu--horizontal .el-menu-item:not(.is-disabled):focus {
  background-color: transparent !important;
}

.el-menu--horizontal>.el-menu-item.is-active {
  border-bottom: none;
}

.el-menu--horizontal.el-menu {
  border-bottom: none; 
  background-image: linear-gradient(to right, var(--el-menu-border-color), var(--el-menu-border-color));
  background-size: calc(100% - 30px) 1px; 
  background-position: 15px bottom;
  background-repeat: no-repeat;
}

.el-menu--horizontal>.el-menu-item {
  border-bottom: none;
}

.icon {
  pointer-events: none;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

.el-dropdown-link {
  cursor: pointer;
  display: flex;
  align-items: center;
}

.el-menu-item:focus,
.el-dropdown-link:focus {
  border: none;
  outline: none;
}

.menu {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 1000;
  backdrop-filter: blur(2px);
  height: 5%;
}

.dark-mode {
  background-color: rgba(32, 33, 36, 0.5);
}

.menu:not(.dark-mode) {
  background-color: rgba(245, 245, 245, 0.5);
}

.menu:not(.dark-mode) .el-dropdown-link {
  color: #000000;
}

.el-menu-item:focus {
  outline: none;
  border: none;
}
.el-menu-item .is-active{
  outline: none !important;
  border: none !important;
}
.flex-grow {
  caret-color: transparent !important;
}
</style>
