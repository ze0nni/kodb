Vue.component("kodb-schema-manager", {
    props: [
        "schema"
    ],
    methods: {
            
    },
    template:
`
<v-dialog>
    <template v-slot:activator="{ on }">
        <v-btn block text v-on="on">
            Edit schema
        </v-btn>
    </template>

    <v-card>
        <v-toolbar flat dark>
            <v-toolbar-title>Schema</v-toolbar-title>
        </v-toolbar>
        <v-tabs vertical>
            <v-tab v-for="t in schema"
                :v-key="t.name"
            >
                <v-icon left>table-large</v-icon>
                {{ t.name }}
            </v-tab>

            <v-tab-item v-for="t in schema"
                :v-key="t.name"
            >
                <kodb-current-schema-manager
                    :table="t"
                >
                </kodb-current-schema-manager>
            </v-tab-item>
            
            <!-- new table -->

            <v-tab>
                <v-icon left>mdi-plus</v-icon>New
            </v-tab>

            <v-tab-item>
                <kodb-new-table-manager
                    :schema="schema"
                >
                </kodb-new-table-manager>
            </v-tab-item>

        </v-tabs>
    </v-card>
</v-dialog>
`
});


Vue.component("kodb-current-schema-manager", {
    props: [
        "table"
    ],
    methods: {
        iconOfType(type) {
            switch (type) {
                case "literal": return "mdi-textbox";
                case "reference": return "mdi-link";
                case "list": return "mdi-view-list";
            }
            return "mdi-help-circle-outline"
        }
    },
    template:
`
<v-card>
    <v-toolbar flat>
        <v-toolbar-title>{{table.name}}</v-toolbar-title>

        <v-spacer></v-spacer>

        <v-btn text color="error">
            Delete
        </v-btn>

    </v-toolbar>
    <v-simple-table>
        <thead>
            <tr>
                <th class="text-left">Column</th>
                <th class="text-left">Type</th>
                <th class="text-left">Options</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="col in table.columns">
                <td>
                    <v-icon left>{{ iconOfType(col.type) }}</v-icon>
                    {{ col.name }}
                </td>
                <td>{{ col.type }}</td>

                <!-- options -->

                <kodb-literal-column-schema
                        v-if="'literal' == col.type"
                ></kodb-literal-column-schema>

                <kodb-ref-column-schema
                        v-else-if="'reference' == col.type"
                ></kodb-ref-column-schema>

                <kodb-list-column-schema
                        v-else-if="'list' == col.type"
                ></kodb-list-column-schema>

                <td  v-else>
                    <v-chip>Unknow type: {{ col.type }}</v-chip>
                </td>

                <!-- /options -->
            </tr>

            <tr>
                <td colspan="3">
                    <v-btn text block>
                        <v-icon left>mdi-plus</v-icon>
                        New
                    </v-btn>
                </td>
            </tr>
        </tbody>
    </v-simple-table>
</v-card>
`
})

Vue.component("kodb-literal-column-schema", {
    props: [
        "col",
        "table"
    ],
    methods: {
            
    },
    template:
`
<td>
    literal
</td>
`
})

Vue.component("kodb-ref-column-schema", {
    props: [
        "col",
        "table"
    ],
    methods: {
            
    },
    template:
`
<td>
    ref
</td>
`
})

Vue.component("kodb-list-column-schema", {
    props: [
        "col",
        "table"
    ],
    methods: {
            
    },
    template:
`
<td>
    list
</td>
`
})

Vue.component("kodb-new-table-manager", {
    props: [
        "schema"
    ],
    template:
`
<v-card>
    <v-tabs>
        <v-tab>
            Literal
        </v-tab>
        <v-tab-item>
            <v-card flat>
                <v-text-field
                >
                </v-text-field>

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
            </v-card>
        </v-tab-item>

        <v-tab>
            Reference
        </v-tab>
        <v-tab-item>
            <v-card>
                <v-select
                    :items="schema"
                >
                </v-select>
            </v-card>
        </v-tab-item>

        <v-tab>
            List
        </v-tab>
        <v-tab-item>
            <v-card>
                <v-select
                    :items="schema"
                >
                </v-select>
            </v-card>
        </v-tab-item>
    </v-tabs>
</v-card>
`
})