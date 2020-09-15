<template>
  <q-page class="page-index">
    <div class="row items-center q-mb-md">
      <q-btn round color="blue-grey" icon="chevron_left" size="sm" @click="backToIndex" />
      <div text-color="blue-grey" class="q-ml-sm text-bold" style="font-size:16px;">Back</div>
    </div>
    <div class="row">
      <div class="title">Create Telemarketer</div>
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
          v-model="email"
          filled
          label="Email Address *"
          type="email"
          class="field"
          :rules="[val => !!val || 'Field is required']"
        />
        <div>
          <q-toggle v-model="isAdmin" checked-icon="check" color="green" unchecked-icon="clear" />Admin
        </div>
        <div class="q-mt-md">
          <q-btn label="Submit" type="submit" color="primary" />
        </div>
      </q-form>
    </div>
  </q-page>
</template>

<script>
import { Dialog } from "quasar";

export default {
  name: "TelemarketerCreate",

  data() {
    return {
      name: "",
      email: "",
      isAdmin: false,
    };
  },

  methods: {
    onSubmit() {
      var vm = this;
      var data_submit = {
        Token: vm.$authService.getToken(),
        Telemarketer: {
          Name: vm.name,
          Email: vm.email,
          IsAdmin: vm.isAdmin,
        },
      };
      this.$axios
        .post("/api/telemarketer/create", data_submit)
        .then(function (response) {
          if (response.data) {
            vm.$router.push({
              name: "telemarketer",
            });
          }
        })
        .catch(function (error) {
          vm.$q.dialog({
            title: "Error Create Telemarketer",
            message: error.response.data.Error,
            cancel: true,
            persistent: true,
          });
        });
      console.log(data_submit);
    },
    backToIndex() {
      this.$router.replace({ name: "telemarketer" });
    },
  },
};
</script>