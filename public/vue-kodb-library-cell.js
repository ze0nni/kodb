Vue.component("kodb-library-cell", {
        props: [
                "libraryName",
                "rowId",
                "column",
                "data"
        ],
        data: function() {
                return {
                        editedValue: ""
                }
        },
        methods: {
                isRowExists(item, colName) {
                        return true
                },
                
                updateValue(rowId, columnId, value) {
                        this.$wsocket.send({
                                "command": "updateValue",
                                "library": this.libraryName,
                                "rowId": this.rowId,
                                "columnId": this.column.value,
                                "value": value
                        })
                }
        },
        template:
`
<div>
        <div
                v-if="isRowExists(data, column.value)"
        >
                <v-edit-dialog
                        @open="editedValue = data[column.value].value"
                        @save="updateValue(rowId, column.value ,editedValue)"
                >
                        {{ data[column.value].value }}
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
</div>
`
});