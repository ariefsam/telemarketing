<template>
  <q-layout>
    <q-page-container class="page-container">
      <q-page class="row justify-center">
        <div class="col-md-6">
          <div class="q-pa-xl">
            <q-btn class="q-mb-md" :to="'/import-customer'">Import Excel Customer</q-btn>
            <q-table title="Treats" :data="data" :columns="columns" row-key="PhoneNumber" />
          </div>
        </div>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script>
export default {
  name: "Login",

  data() {
    return {
      columns: [
        {
          name: "Name",
          label: "Name",
          field: "Name",
          align: "left",
        },
        {
          align: "left",
          name: "PhoneNumber",
          label: "Phone Number",
          field: "PhoneNumber",
        },
        {
          align: "left",
          label: "Status",
          field: "Status",
          name: "Status",
        },
      ],
      data: [
        {
          Name: "Arief",
          PhoneNumber: "09343",
          Status: "Not Set",
        },
        {
          Name: "Ay",
        },
      ],
    };
  },
};
</script>