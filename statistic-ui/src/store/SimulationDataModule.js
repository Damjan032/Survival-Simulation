import Vue from "vue";
import axios from "axios";

const SimulationDataModule = {
    namespaced: true,
    state: {
        currentId: null,
        fullData: null,
        mainDates: []
    },
    mutations: {
        setFullData(state, fullData) {
            Vue.set(state, "fullData", fullData);
        },
        setMainDates(state, mainDates) {
            Vue.set(state, "mainDates", mainDates);
        },
    },
    actions: {
        async getFullDataById({ commit }, id) {
            try {
                axios
                    .get("fullData/"+id)
                    .then((res) => {
                        if (res.data) {
                            commit("setFullData", res.data);
                        }
                    })
                    .catch((err) => {
                        console.log(err);
                    });
            } catch (err) {
                console.log(err);
            }
        },
        async getFullFirstData({ commit }) {
            try {
                axios
                    .get("firstData")
                    .then((res) => {
                        if (res.data) {
                            commit("setFullData", res.data);
                        }
                    })
                    .catch((err) => {
                        console.log(err);
                    });
            } catch (err) {
                console.log(err);
            }
        },
        async getMainDates({ commit }) {
            try {
                axios
                    .get("mainData")
                    .then((res) => {
                        if (res.data) {
                            commit("setMainDates", res.data);
                        }
                    })
                    .catch((err) => {
                        console.log(err);
                    });
            } catch (err) {
                console.log(err);
            }
        },
    },
};
export default SimulationDataModule;
