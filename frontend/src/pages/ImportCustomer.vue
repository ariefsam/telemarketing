<template>
  <q-layout>
    <q-page-container class="page-container">
      <q-page class="row justify-center">
        <div class="col-md-6">
          <div class="q-pa-xl">
            Upload:
            <q-uploader :url="uploadUrl" style="max-width: 300px" auto-upload field-name="file" />
          </div>
        </div>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script>
export default {
  name: "ImportCustomer",

  data() {
    return {
      uploadUrl: "",
    };
  },
  mounted() {
    var vm = this;
    var token = vm.$authService.getToken();
    vm.uploadUrl = "/api/upload?token=" + token;
  },
};
</script>