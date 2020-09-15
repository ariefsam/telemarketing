<template>
  <q-page class="page-index">
    <div class="row items-center q-mb-md">
      <q-btn round color="blue-grey" icon="chevron_left" size="sm" @click="backToIndex" />
      <div text-color="blue-grey" class="q-ml-sm text-bold" style="font-size:16px;">Back</div>
    </div>
    <div class="row">
      <div class="title">Edit Telemarketer</div>
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
        <div class="q-mb-sm">
          <q-toggle v-model="isAdmin" checked-icon="check" color="green" unchecked-icon="clear" />Admin
        </div>
        <div class="q-mb-sm">
          Weekly Target
        </div>
        <div class="row q-col-gutter-sm">
          <q-input
            color="grey-8"
            v-model.number="callTarget"
            filled
            label="Call Target *"
            type="number"
            class="field col-md-6"
            :rules="[val => !!val || 'Field is required']"
          />
          <q-input
            color="grey-8"
            v-model.number="closingTarget"
            filled
            label="Closing Target *"
            type="number"
            class="field col-md-6"
            :rules="[val => !!val || 'Field is required']"
          />
        </div>
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
  name: "TelemarketerEdit",

  data() {
    return {
      name: "",
      email: "",
      isAdmin: false,
      callTarget: null,
      closingTarget: null,
    };
  },

  mounted(){
    // this.name = "Agung"
    // this.email = "agung@fsn.co.id"
    // this.isAdmin = true
    // this.callTarget = 10
    // this.closingTarget = 10
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
          WeeklyTarget: {
            Call: vm.callTarget,
            Closing: vm.closingTarget,
          },
        },
      };
      // this.$axios
      //   .post("/api/telemarketer/create", data_submit)
      //   .then(function (response) {
      //     if (response.data) {
      //       vm.$router.push({
      //         name: "telemarketer",
      //       });
      //     }
      //   })
      //   .catch(function (error) {
      //     vm.$q.dialog({
      //       title: "Error Edit Telemarketer",
      //       message: error.response.data.Error,
      //       cancel: true,
      //       persistent: true,
      //     });
      //   });
      console.log(data_submit);
    },
    backToIndex() {
      this.$router.replace({ name: "telemarketer" });
    },
  },
};
</script>