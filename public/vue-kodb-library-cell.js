Vue.component("kodb-library-cell", {
        props: [
                "schema",
                "libraryName",

                "row",
                "field",

                "expandRow",
                "isExpanded"
        ],
        data: function() {
                return {
                        
                }
        },
        computed: {
                cellData() {
                        if (null == this.row || null == this.row.data || null == this.row.data[this.field.id]) {
                                return {
                                        exists: false
                                }
                        }
                        return this.row.data[this.field.id];
                }
        },
        methods: {           
                updateValue(value) {
                        this.$wsocket.send({
                                "command": "updateValue",
                                "library": this.libraryName,
                                "rowId": this.row.rowId,
                                "columnId": this.columnId,
                                "value": value
                        })
                },
        },
        template:
`
<v-row>
        <kodb-library-value-cell 
                v-if="'value' == field.kind"

                v-on:update="updateValue"

                :cellData="cellData"
                :field="field"
        ></kodb-library-value-cell>

        <kodb-library-reference-cell 
                v-else-if="'reference' == field.kind"

                v-on:update="updateValue"
                
                :schema="schema"
                :column="getColumnData(libraryName, columnId)"
                :rowData="rowData"
                :cellData="cellData"

        ></kodb-library-reference-cell>

        <kodb-library-list-cell 
                v-else-if="'list' == field.kind"

                v-on:update="updateValue"

                :schema="schema"
                :libraryName="libraryName"
                :column="getColumnData(libraryName, columnId)"
                :rowData="rowData"
                :cellData="cellData"

                :expandRow="expandRow"
                :isExpanded="isExpanded"

        ></kodb-library-list-cell>

        <v-chip
                v-else
                color="error"
        >
                Unknown type: {{ field.kind }}
        </v-chip>
</v-row>
`
});

Vue.component("kodb-library-value-cell", {
        props: [
                "cellData",
                "field"
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
        <v-edit-dialog
                @open="editedValue = cellData.value"
                @save="$emit('update', editedValue)"
        >
                <div v-if="cellData.exists && cellData.value">
                        {{ cellData.value }}
                </div>
                <div v-else-if="cellData.exists">
                        <v-icon small>mdi-cursor-text</v-icon>
                </div>
                <v-chip v-else
                        x-small
                >
                        nil
                </v-chip>

                <template v-slot:input>
                        <v-text-field
                                v-model="editedValue"
                                label="Edit"
                                single-line
                        ></v-text-field>
                </template>
        </v-edit-dialog>
</div>
`
})

Vue.component("kodb-library-reference-cell", {
        props: [
                "schema",
                "column",
                "rowData",
                "cellData"
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
        :items="mapItems(schema.rowsMap[column.reference], column)"
>
</v-select>
`
})

Vue.component("kodb-library-list-cell", {
        props: [
                "schema",
                "libraryName",
                "column",
                "rowData",
                "cellData",

                "expandRow",
                "isExpanded"
        ],
        data: function() {
                return {
                        selectedItem: this.cellData.value
                }
        },
        methods: {
                filterItems(items) {
                        const libraryName = this.libraryName
                        const rowId = this.rowData.rowId
                        const columnId = this.column.id
                        return (items || [])
                                .filter(r => this.isParentOf(
                                        libraryName,
                                        rowId,
                                        columnId,
                                        r
                                ))
                }
        },
        watch: {
                selectedItem(value) {
                        this.$emit('update', value)
                }
        },
        template:
`
<v-btn small text
        v-on:click="expandRow"
>
        [ {{ column.name }}({{ filterItems(schema.rowsMap[column.reference]).length }}) ]

        <v-icon v-if="!isExpanded">
                mdi-chevron-down
        </v-icon>
        <v-icon v-if="isExpanded">
                mdi-chevron-up
        </v-icon>
</v-btn>
`
})