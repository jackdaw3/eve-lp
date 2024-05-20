import { createStore } from 'vuex'

interface State {
  regionId: string,
  form: {
    materialPrice: string;
    productPrice: string;
    days: string;
    scope: string;
    tax: number;
  };
  tableForm: {
    calculator: string;
    layout: string;
    highlight: string;
  };
  corporationId: number;
  itemVisible: boolean;
  itemDetail: any[];
  calculatorVisible: boolean;
  selectedItems: any[];
}

export default createStore<State>({
  state: {
    regionId: "10000002",
    form: {
      materialPrice: "sell",
      productPrice: "buy",
      days: "7",
      scope: "0.05",
      tax: 0
    },
    tableForm: {
      calculator: "false",
      layout: "auto",
      highlight: "false",
    },
    corporationId: 0,
    itemVisible: false,
    itemDetail: [],
    calculatorVisible: false,
    selectedItems: [],
  },
  mutations: {
    setRegionId(state, value: string) {
        state.regionId = value;
    },
    setForm(state, value: State['form']) {
      state.form = value;
    },
    setTableForm(state, value: State['tableForm']) {
      state.tableForm = value;
    },
    setCorporationId(state, value: number) {
      state.corporationId = value;
    },
    setItemVisible(state, value: boolean) {
      state.itemVisible = value;
    },
    setItemDetail(state, value: any[]) {
      state.itemDetail = value;
    },
    setCalculatorVisible(state, value: boolean) {
      state.calculatorVisible = value;
    },
    setSelectedItems(state, value: any[]) {
      state.selectedItems = value;
    },
  },
  actions: {
    setRegionId({ commit }, value: string) {
        commit("setRegionId", value);
    },
    setForm({ commit }, value: State['form']) {
      commit("setForm", value);
    },
    setTableForm({ commit }, value: State['tableForm']) {
      commit("setTableForm", value);
    },
    setCorporationId({ commit }, value: string) {
      commit("setCorporationId", value);
    },
    setQuickBarSelected({ commit }, value: any[]) {
      commit("setQuickBarSelected", value);
    },
    setItemVisible({ commit }, value: boolean) {
      commit("setItemVisible", value);
    },
    setItemDetail({ commit }, value: any[]) {
      commit("setItemDetail", value);
    },
    setCalculatorVisible({ commit }, value: boolean) {
      commit("setCalculatorVisible", value);
    },
    setSelectedItems({ commit }, value: any[]) {
      commit("setSelectedItems", value);
    },
  },
})
