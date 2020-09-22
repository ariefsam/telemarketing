<template>
  <q-page class="page-index">
    <div class="row items-center q-mb-md">
      <q-btn round color="blue-grey" icon="chevron_left" size="sm" @click="backToIndex" />
      <div text-color="blue-grey" class="q-ml-sm text-bold" style="font-size:16px;">Back</div>
    </div>
    <div class="row">
      <div class="title">Add Customer</div>
    </div>
    <div class="row">
      <q-form @submit="onSubmit" class="col-md-6">
        <q-input
          color="grey-8"
          v-model="name"
          filled
          label="Name *"
          type="text"
          class="field"
          :rules="[val => !!val || 'Field is required']"
        />
        <q-input
          color="grey-8"
          v-model="phoneNumber"
          filled
          label="Phone Number *"
          class="field"
          :rules="[val => !!val || 'Field is required']"
        />
        <div class="q-mt-sm">
          <q-btn label="Submit" type="submit" color="primary" />
        </div>
      </q-form>
    </div>
  </q-page>
</template>

<script>
import { Dialog } from "quasar";

export default {
  name: "CustomerAdd",

  data() {
    return {
      name: "",
      phoneNumber: "",
    };
  },

  methods: {
    onSubmit() {
      var vm = this;
      var user = vm.$authService.getUser();
      var data_submit = {
        Token: vm.$authService.getToken(),
        Customer: {
          Name: vm.name,
          PhoneNumber: vm.phoneNumber,
          DataSource: user.ID,
          CreatedBy: user.Name,
        },
      };
      this.$axios
        .post("/api/customer/create", data_submit)
        .then(function (response) {
          if (response.data) {
            vm.$router.replace({
              name: "customer"
            });
          }
        })
        .catch(function (error) {
          vm.$q.dialog({
            title: "Error Add Customer",
            message: error.response.data.Error,
            persistent: false,
          });
        });
      console.log(data_submit);
    },
    backToIndex() {
      this.$router.replace({ name: "customer" });
    },
  },
};
</script>