<template>
  <q-page class="page-index">
    <div class="row justify-between">
      <div class="title">Telemarketer</div>
      <div>
        <q-btn
          color="light-green-9"
          text-color="white"
          icon="add"
          label="Add"
          class="q-pr-sm"
          :to="{ name: 'telemarketer-create' }"
        />
      </div>
    </div>
    <div class="table-container">
      <q-table
        :data="telemarketers"
        :columns="telemarketerDataColumns"
        row-key="email"
        :filter="telemarketerDataFilter"
        :visible-columns="telemarketerDataVisible"
        :pagination.sync="telemarketerDataPagination"
        @row-click="onRowClick"
      >
        <template v-slot:top-right>
          <q-input color="grey-8" dense debounce="300" v-model="telemarketerDataFilter" placeholder="Search">
            <template v-slot:append>
              <q-icon name="search" />
            </template>
          </q-input>
        </template>
        <template v-slot:top-left>
          <div>
            <span class="q-pr-sm">Total Telemarketer</span>
            <span class="text-bold" style="font-size: 14px">{{telemarketers.length}}</span>
          </div>
        </template>
        <q-td slot="body-cell-isAdmin" slot-scope="props" :props="props">
          <div v-if="props.value == true">Yes</div>
          <div v-else>No</div>
        </q-td>
      </q-table>
    </div>
  </q-page>
</template>

<script>
export default {
  name: 'TelemarketerIndex',

  data() {
    return {
      telemarketers: [],
      telemarketerDataColumns: [
        { name: 'name', label: 'NAME', align: 'left', field: 'Name', sortable: true },
        { name: 'email', label: 'EMAIL', align: 'left', field: 'Email', sortable: true },
        { name: 'isAdmin', label: 'Admin', align: 'left', field: 'IsAdmin', sortable: true },
      ],
      telemarketerDataFilter: "",
      telemarketerDataVisible: ['name', 'email', 'isAdmin'],
      telemarketerDataPagination: {
        rowsPerPage: 5 // current rows per page being displayed
      },
    }
  },

  mounted() {
    var vm = this;
    var data_submit = {
      Token: vm.$authService.getToken(),
      Limit: 10000,
    }
    this.$axios
      .post("/api/telemarketer", data_submit)
      .then(function (response) {
        if (response.data) {
          vm.telemarketers = response.data.Telemarketers
        }
      })
    // dummy
    vm.telemarketers = [
      {
        Name: "Arief Hidayatulloh",
        Email: "arief@fsn.co.id",
        IsAdmin: true,
      },
      {
        Name: "Agung Kurniawan",
        Email: "agung@fsn.co.id",
        IsAdmin: true,
      }
    ]
  },

  methods: {
    onRowClick(evt, row) {
      console.log("clicked on", row.Email);
    },
  }
}
</script>