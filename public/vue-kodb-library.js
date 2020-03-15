Vue.component("kodb-library", {
        props: [
                "schema",
                "libraryName",
        ],
        data: function() {
                return {
                        multiSelect: false,
                        selectedRows:[],
                        expandedLibraryName: null
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
                                "library": this.libraryName
                        })
                },

                deleteSelectedRows() {
                        for (let row of this.selectedRows) {
                                this.$wsocket.send({
                                        "command": "deleteRow",
                                        "library": this.libraryName,
                                        "rowId": row.rowId
                                })
                        }
                        this.selectedRows = []
                },
        },
        template:
`
<v-data-table
        :headers="mapColumns(schema.map[libraryName].columns)"
        :items="schema.rowsMap[libraryName]"
        :items-per-page="10"
        item-key="rowId"

        v-model="selectedRows"

        dense
        show-select
        :single-select="!multiSelect"
>       
        <template v-slot:item="{ item,headers,select,isSelected,expand,isExpanded }">
        
                <tr v-on:click="select(!isSelected)">
                        <td v-for="col in headers"
                        >
                                <v-icon
                                        v-if="col.value == 'data-table-select'"
                                >
                                        {{ isSelected ? "mdi-check-box-outline" : "mdi-checkbox-blank-outline" }}
                                </v-icon>

                                <kodb-library-cell
                                        v-else

                                        :schema="schema"
                                        :libraryName="libraryName"
                                        :rowId="item.rowId"
                                        :columnId="col.id"

                                        :rowData="item"
                                        :cellData="item.data[col.value]"
                                        
                                        :expandRow="() => {
                                                if (expandedLibraryName != col.reference) {
                                                        expandedLibraryName = col.reference
                                                        expand(true);
                                                } else {
                                                        expand(!isExpanded);
                                                }
                                        }"
                                        :isExpanded="expandedLibraryName == col.reference && isExpanded"
                                >
                                </kodb-library-cell>
                        </td>
                </tr>
        </template>

        <template v-slot:expanded-item="{ item, headers }">
                <td></td>
                <td :colspan="headers.length-1"
                >
                        <kodb-library-expanded
                                depth=1
                                :schema="schema"
                                :libraryName="expandedLibraryName"
                                
                                :parentLibraryName="libraryName"
                                :parentRowId="item.rowId"
                        >
                        </kodb-library-expanded>
                </td>
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