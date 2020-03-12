Vue.component("kodb-library-cell", {
        props: [
                "libraryName",
                "rowId",
                "column",
                "data",
                "librarisData"
        ],
        data: function() {
                return {
                        
                }
        },
        methods: {                
                updateValue(value) {
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
<v-row>
        <kodb-library-literal-cell 
                v-if="'literal' == column.type"

                v-on:update="updateValue"
                :column="column"
                :data="data"
                :cellData="data[column.value]"
        ></kodb-library-literal-cell>

        <kodb-library-reference-cell 
                v-else-if="'reference' == column.type"

                v-on:update="updateValue"
                :libraryName="libraryName"
                :rowId="rowId"
                :column="column"
                :data="data"
                :cellData="data[column.value]"
                :librarisData="librarisData"

        ></kodb-library-reference-cell>

        <kodb-library-list-cell 
                v-else-if="'list' == column.type"

                v-on:update="updateValue"
                :libraryName="libraryName"
                :rowId="rowId"
                :column="column"
                :data="data"
                :cellData="data[column.value]"
                :librarisData="librarisData"

        ></kodb-library-list-cell>

        <v-chip
                v-else
                color="error"
        >
                Unknown type: {{ column.type }}
        </v-chip>
</v-row>
`
});

Vue.component("kodb-library-literal-cell", {
        props: [
                "column",
                "data",
                "cellData"
        ],
        data() {
                return {
                        editedValue: ""
                }
        },
        methods: {
                isRowExists(item, colName) {
                        return item[colName]
                                && item[colName].exists
                }
        },
        template:
`
<div
        <div v-if="isRowExists(data, column.value)">
                <v-edit-dialog
                        @open="editedValue = cellData.value"
                        @save="$emit('update', editedValue)"
                >
                        {{ cellData.value }}
                        <template v-slot:input>
                                <v-text-field
                                        v-model="editedValue"
                                        label="Edit"
                                        single-line
                                ></v-text-field>
                        </template>
                </v-edit-dialog>
        </div>
        <v-chip v-else>
                nil
        </v-chip>
</div>
`
})

Vue.component("kodb-library-reference-cell", {
        props: [
                "libraryName",
                "rowId",
                "column",
                "data",
                "cellData",
                "librarisData"
        ],
        data: function() {
                return {
                        selectedItem: this.cellData.value
                }
        },
        methods: {
                mapItems(items, column) {
                        return (items || [])
                        .map(function(r) {
                                return {
                                        text: r.rowId,
                                        value: r.rowId
                                }
                        })
                }
        },
        watch: {
                selectedItem(value) {
                        this.$emit('update', value)
                }
        },
        template:
`
<v-select
        v-model="selectedItem"

        :error-messages="cellData.error"
        :items="mapItems(librarisData[column.reference], column)"
>
</v-select>
`
})

Vue.component("kodb-library-list-cell", {
        props: [
                "libraryName",
                "rowId",
                "column",
                "data",
                "cellData",
                "librarisData"
        ],
        data: function() {
                return {
                        selectedItem: this.cellData.value
                }
        },
        methods: {
                filterItems(items) {
                        return (items || [])
                                .filter(r => {
                                        const parent = r.data.parent
                                        return parent && parent.value == this.rowId
                                })
                }
        },
        watch: {
                selectedItem(value) {
                        this.$emit('update', value)
                }
        },
        template:
`
<div>
        <table>
                <tr v-for="r in filterItems(librarisData[column.reference])"
                >
                        <td>{{ r.rowId }}</td>
                </tr>
        </table>
</div>
`
})