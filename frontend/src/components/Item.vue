<template>
    <div class="Item">
      <el-dialog v-model="itemVisible" v-if="itemVisible" class="custom-dialog">
        <template #header>
          <div style="display: flex; align-items: center; max-height: 50%;">
            <el-image style="margin-right: 8px; height: 52px;" :src="getIcon" v-if="itemVisible">
              <template #error>
                <div class="image-slot" style="font-size: 52px;">
                  <i class="el-icon-picture-outline"></i>
                </div>
              </template>
            </el-image>
            <div>
              <div style="font-size: 18px;">{{ itemDetail.ItemName }}</div>
              <br class="short-br">
              <div style="font-size: 18px;">{{ itemDetail.Volume }}</div>
            </div>
          </div>
        </template>
        <el-divider></el-divider>
        <span v-html="removeHtmlTags(itemDetail.Description)"></span>
      </el-dialog>
    </div>
  </template>
  
  <script setup lang="ts">
  import { computed } from 'vue';
  import { useStore } from 'vuex';
  import { ICON_SERVER } from '~/constants';
  
  const store = useStore();
  
  const itemVisible = computed({
    get: () => store.state.itemVisible,
    set: (value) => store.commit('setItemVisible', value)
  });
  const itemDetail = computed(() => store.state.itemDetail);
  const getIcon = computed(() => ICON_SERVER + "Type/" + itemDetail.value.ItemId + "_64.png");
  
  function removeHtmlTags(str: string): string {
    if (str) {
      str = str.replace(/(\n+)/g, '<br><br>');
      str = str.replace(/(<br>\s*){3,}/g, '<br><br>');
      return str.replace(/<[^>]+>/g, (match) => {
        return match === '<br>' || match === '<br/>' ? match : '';
      });
    }
    return "";
  }
  </script>
  
  <style>
  .custom-dialog .el-dialog__header {
    padding: 10px 10px;
    max-height: 28px;
  }
  
  .short-br {
    display: block;
    content: "";
    margin-top: 6px;
  }
  
  .el-image {
    caret-color: transparent !important;
  }
  </style>
  