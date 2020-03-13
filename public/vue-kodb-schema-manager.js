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
            
    },
    template:
`
<v-simple-table>
    <thead>
        <tr>
            <th class="text-left">Column</th>
            <th class="text-left">Type</th>
        </tr>
    </thead>
    <tbody>
        <tr v-for="col in table.columns">
          <td>{{ col.name }}</td>
          <td>{{ col.type }}</td>
        </tr>
      </tbody>
</v-simple-table>
`
})