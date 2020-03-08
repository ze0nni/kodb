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
        <v-avatar v-if="data[column.value].error"
                color="red"
                size="14"
        >
                !
        </v-avatar>

        <kodb-library-literal-cell 
                v-if="'literal' == column.type"

                v-on:update="updateValue"
                :column="column"
                :data="data"
        ></kodb-library-literal-cell>

        <kodb-library-reference-cell 
                v-else-if="'reference' == column.type"

                v-on:update="updateValue"
                :libraryName="libraryName"
                :rowId="rowId"
                :column="column"
                :data="data"
                :librarisData="librarisData"

        ></kodb-library-reference-cell>

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
                "data"
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
                        @open="editedValue = data[column.value].value"
                        @save="$emit('update', editedValue)"
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
                "librarisData"
        ],
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
        template:
`
<v-select
        :items="mapItems(librarisData[column.reference], column)"
>
</v-select>
`
})