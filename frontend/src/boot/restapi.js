import Vue from 'vue'
var vm = Vue.prototype;
var restapi = {
    callCustomer(data_submit, callbackSuccess) {
        Vue.prototype.$axios
            .post("/api/customer/call", data_submit)
            .then(function (response) {
                if (response.data) {
                    vm.$q
                        .dialog({
                            title: "Success",
                            message: "Success",
                        })
                        .onOk(() => {
                            // console.log('OK')
                        })
                        .onCancel(() => {
                            // console.log('Cancel')
                        })
                        .onDismiss(() => {
                            callbackSuccess();
                        });

                }
            })
            .catch(function (error) {
                console.log(error);
            });
    },
}

Vue.prototype.$restapi = restapi