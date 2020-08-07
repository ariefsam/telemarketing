<template>
  <q-layout>
    <q-page-container class="page-container">
      <q-page class="row justify-center">
        <div class="col-md-6">
          <div class="q-pa-xl">
            <q-table title="Telemarketers" :data="data" :columns="columns" row-key="PhoneNumber" />
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
          name: "Email",
          label: "Email",
          field: "Email",
        },
      ],
      data: [
        {
          Name: "Arief",
          Email:"arief@gmail.com"
        },
      ],
    };
  },
};
</script>