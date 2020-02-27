Vue.component("kodb-library", {
        props: [
                "librarySchema"
        ],
        data: function() {
                return {
                        editedValue: "",
                        rows:[],
                }
        },
        methods: {
                mapColumns(columns) {
                        return columns.map(col => {
                                return {
                                        text: col.name,
                                        value: col.id
                                }
                        }).concat([
                                {
                                        text: "actions",
                                        actions: true
                                }
                        ])
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

                deleteRow(rowId) {
                        this.$wsocket.send({
                                "command": "deleteRow",
                                "library": this.librarySchema.name,
                                "rowId": rowId
                        })
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
>
        <template v-slot:item="{ item,headers }">
                <tr>
                        <td v-for="col in headers"
                        >
                                <div    v-if="col.actions">
                                        <v-icon
                                                v-on:click="deleteRow(item.rowId)"
                                        >
                                                mdi-delete
                                        </v-icon>
                                </div>
                                <div
                                        v-else-if="isRowExists(item, col.value)"
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
                        <v-spacer></v-spacer>
                        <v-btn v-on:click="newRow">New row</v-btn>
                </v-toolbar>
        </template>
</v-data-table>
`
});