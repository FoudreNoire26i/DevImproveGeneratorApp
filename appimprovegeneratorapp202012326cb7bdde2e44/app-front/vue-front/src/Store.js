import Vue from "vue";
import VueX from "vuex"

Vue.use(VueX);

const state = {
  products:[],
  adresse:[],
  searched: {question:"",field:""},
  requests: new Map(),
  requestsLength: 0,
  selected: {},
}

const mutations = {
  ADD_PRODUCT: (state,name) => {
    state.products.push(name)
  },
  ADD_ADRESSE:(state,name) => {
    state.adresse.push(name)
  },
  SUPP_PRODUCT: (state,name) => {
    state.products.splice(name,1)
  },
  SET_SEARCHED: (state, searched) => {
    state.searched = searched
  },
  RESET_SEARCHED: (state) => {
    state.searched = {question:"",field:""};
  },
  PUSH_REQUEST: (state, item) => {
    if(!state.requests.has(item.question)){
      state.requests.set(item.question, [item])
    }
      
    else state.requests.get(item.question).push(item)
  },
  REMOVE_REQUEST: (state, idItem) => {
    if(state.requests.has(item.question)) {
      const index = state.requests.get(item.question).indexOf(item);
      if (index > -1) {
        state.requests.get(item.question).splice(index, 1);
      }
    }
    if(state.requests.get(item.question).length === 0) state.requests.delete(item.question)
  },
  FIND_ITEM_FROM_ID: (state, idItem) => {
    
  },

  SET_REQUESTS_LENGTH: (state) => {
    state.requestsLength = 0;
    for (var val of state.requests.values()) {
      console.log(val);
      state.requestsLength += val.length;
    }
  } ,
  SET_SELECTED: (state, item) => {
    state.selected = item;
  }
}

const getters = {
  products: state => state.products,
  adresse: state => state.adresse,
  searched: state => state.searched,
  requests: state => state.requests,
  requestsLength: state => state.requestsLength,
  selected: state => state.selected,
}
export default new VueX.Store({
  state: state,
  mutations: mutations,
  getters: getters,
  actions: {}
})

