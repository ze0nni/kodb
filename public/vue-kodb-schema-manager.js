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
                <v-icon left>mdi-account</v-icon>
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
      </tbody>
</v-simple-table>
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