Vue.component("kodb-library", {
        props: [
                "librarySchema",
                "rows",
                "librarisData"
        ],
        data: function() {
                return {
                        multiSelect: false,
                        selectedRows:[],
                }
        },
        methods: {
                mapColumns(columns) {
                        return columns.map(col => {
                                return Object.assign(
                                        {},
                                        col,
                                        {
                                                text: col.name,
                                                value: col.id
                                        }
                                )
                        })
                },

                newRow() {
                        this.$wsocket.send({
                                "command": "newRow",
                                "library": this.librarySchema.name
                        })
                },

                deleteSelectedRows() {
                        for (let row of this.selectedRows) {
                                this.$wsocket.send({
                                        "command": "deleteRow",
                                        "library": this.librarySchema.name,
                                        "rowId": row.rowId
                                })
                        }
                        this.selectedRows = []
                },
        },
        template:
`
<v-data-table
        :headers="mapColumns(librarySchema.columns)"
        :items="rows"
        :items-per-page="10"
        item-key="rowId"

        v-model="selectedRows"

        show-select
        :single-select="!multiSelect"
>       
        <template v-slot:item="{ item,headers,select,isSelected }">
        
                <tr v-on:click="select(!isSelected)">
                        <td v-for="col in headers"
                        >
                                <v-checkbox
                                        v-if="col.value == 'data-table-select'"
                                        v-bind:value="isSelected">
                                </v-checkbox>

                                <kodb-library-cell
                                        v-else
                                        :libraryName="librarySchema.name"
                                        :rowId="item.rowId"
                                        :column="col"
                                        :data="item.data"
                                        :librarisData="librarisData"
                                >
                                </kodb-library-cell>
                        </td>
                </tr>
        </template>

        <template v-slot:top>
                <v-toolbar flat>
                        <v-switch v-model="multiSelect" label="Multi select" />

                        <v-spacer></v-spacer>

                        <v-btn v-on:click="newRow"
                                icon color="primary"
                        >
                                <v-icon>mdi-plus</v-icon>
                        </v-btn>
                        <v-btn v-on:click="deleteSelectedRows"
                                :disabled="selectedRows.length == 0"
                                icon color="error"
                        >
                                <v-icon>mdi-delete</v-icon>
                        </v-btn>

                        <v-btn :disabled="selectedRows.length == 0"
                                icon color="primary"
                        >
                                <v-icon>mdi-arrow-up</v-icon>
                        </v-btn>
                        <v-btn :disabled="selectedRows.length == 0"
                                icon color="primary"
                        >
                                <v-icon>mdi-arrow-down</v-icon>
                        </v-btn>
                </v-toolbar>
        </template>
</v-data-table>
`
});