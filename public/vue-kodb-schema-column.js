(function() {

Vue.component("vue-kodb-schema-literal-column", {
        props: [
                "column",

                "confirm"
        ],
        data() {
                return {
                        columnName: this.column.name || ""
                }
        },
        methods: {
                submit() {
                        this.confirm({
                                name: this.columnName,
                                type: "literal"
                        })
                }
        },
        template:
`
<v-card>
        <v-col>
                <v-text-field
                        v-model="columnName"
                        label="Column name"
                ></v-text-field>

                <v-radio-group>
                <v-radio
                        label="String"
                ></v-radio>
                <v-radio
                        label="Text"
                ></v-radio>
                <v-radio
                        label="Int"
                ></v-radio>
                <v-radio
                        label="Float"
                ></v-radio>
                <v-radio
                        label="Option"
                ></v-radio>
                <v-radio
                        label="Set"
                ></v-radio>
                </v-radio-group>
        </v-col>

        <v-card-actions>
                <v-btn text block
                        v-on:click="submit"
                >OK</v-btn>
        </v-card-actions>
</v-card>
`})

//
const ColumnsToItemsMixin = {
        methods: {
                toItems(items) {
                        return items.map(x => ({
                                "text": x.name,
                                "value": x.name
                        }))
                }
        }
}

Vue.component("vue-kodb-schema-ref-column", {
        mixins: [ColumnsToItemsMixin],
        props: [
                "schema",
                "column",

                "confirm"
        ],
        data() {
                return {
                        columnName: this.column.name || "",
                        "selectedLibrary": this.column.reference || null,
                        "displayColumn": this.column["meta:displayColumn"],
                }
        },
        methods: {
                toColumnItems(library) {
                        if (null == library) {
                                return []
                        }
                        return library.columns.map(c => ({
                                "text": c.name,
                                "value": c.id
                        }))
                },
                submit() {
                        this.confirm({
                                library: this.libraryName,
                                name: this.columnName,
                                type: "reference",
                                "ref": this.selectedLibrary,
                                
                                "meta:displayColumn": this.displayColumn
                        })
                }
        },
        template:
`
<v-card>
        <v-col>
                <v-text-field
                        v-model="columnName"
                        label="Column name"
                ></v-text-field>

                <v-select
                        v-model="selectedLibrary"
                        :items="toItems(schema.list)"

                        label="Library"
                >
                </v-select>

                <v-select
                        v-model="displayColumn"
                        :items="toColumnItems(schema.map[selectedLibrary])"

                        label="Column to display"
                >
                </v-select>
        </v-col>

        <v-card-actions>
                <v-btn text block
                        v-on:click="submit"
                >OK</v-btn>
        </v-card-actions>
</v-card>
`})

Vue.component("vue-kodb-schema-list-column", {
        mixins: [ColumnsToItemsMixin],
        props: [
                "schema",

                "confirm"
        ],
        data() {
                return {
                        columnName: "",
                        selectedLibrary: null
                }
        },
        methods: {
                submit() {
                        this.confirm({
                                name: this.columnName,
                                type: "list",
                                "ref": this.selectedLibrary
                        })
                }
        },
        template:
`
<v-card>
        <v-col>
                <v-text-field
                        v-model="columnName"
                        label="Column name"
                ></v-text-field>

                <v-select
                        v-model="selectedLibrary"
                        :items="toItems(schema.list)"
                >
                </v-select>

                <v-select
                        v-model="selectedLibrary"
                        :items="toItems(schema.list)"
                >
                </v-select>
        </v-col>

        <v-card-actions>
                <v-btn text block
                        v-on:click="submit"
                >OK</v-btn>
        </v-card-actions>
</v-card>
`})


})()