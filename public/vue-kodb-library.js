Vue.component("kodb-library", {
        props: [
                "librarySchema"
        ],
        data: function() {
                return {
                        multiSelect: false,
                        editedValue: "",
                        rows:[],
                        selectedRows:[],
                }
        },
        methods: {
                mapColumns(columns) {
                        return columns.map(col => {
                                return {
                                        text: col.name,
                                        value: col.id
                                }
                        })
                },
                isRowExists(item, colName) {
                        return item
                                && item.data
                                && item.data[colName]
                                && item.data[colName].exists
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

                updateValue(rowId, columnId, value) {
                        this.$wsocket.send({
                                "command": "updateValue",
                                "library": this.librarySchema.name,
                                "rowId": rowId,
                                "columnId": columnId,
                                "value": value
                        })
                }
        },
        webSockets: {
                connected() {
                        this.$wsocket.send({
                                "command": "getLibraryRows",
                                "library": this.librarySchema.name
                        })
                },
                command: {
                        setLibraryRows(msg) {
                                if (msg.library != this.librarySchema.name) {
                                        return
                                }
                                this.rows = msg.rows
                        },
                        newRow(msg) {
                                if (msg.library != this.librarySchema.name) {
                                        return
                                }
                                this.rows.push({
                                        "rowId": msg.rowId,
                                        "data": {}
                                })
                        },
                        deleteRow(msg) {
                                if (msg.library != this.librarySchema.name) {
                                        return
                                }
                                const rowId = msg.rowId
                                constRowIndex = this.rows.findIndex(x => x.rowId == rowId)
                                if (-1 != constRowIndex) {
                                        this.rows.splice(constRowIndex, 1)
                                }
                        }
                }
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
        <template v-slot:item="{ item,headers }">
                <tr>
                        <td v-for="col in headers"
                        >
                                <div
                                        v-if="isRowExists(item, col.value)"
                                >
                                        <v-edit-dialog
                                                @open="editedValue = item.data[col.value].value"
                                                @save="updateValue(item.rowId, col.value ,editedValue)"
                                        >
                                                {{ item.data[col.value].value }}
                                                <template v-slot:input>
                                                        <v-text-field
                                                                v-model="editedValue"
                                                                label="Edit"
                                                                single-line
                                                        ></v-text-field>
                                                </template>
                                        </v-edit-dialog>
                                </div>
                                <v-chip 
                                        v-else
                                        color="red"
                                >
                                        NIL
                                </v-chip>
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