Vue.component("vue-kodb-schema-new-column", {
        props: [
            "schema",
            "libraryName"
        ],
        data() {
            return {
                "dialog": false
            }
        },
        methods: {
            confirm(msg) {
                msg['command'] = "newColumn"
                this.$wsocket.send(msg)
            }
        },
        template:
    `
<v-dialog
        v-model="dialog"
>
    <template v-slot:activator="{ on }">
        <v-btn text block  v-on="on">
            <v-icon left>mdi-plus</v-icon>
            New
        </v-btn>
    </template>

    <v-card v-if="dialog">
        <v-toolbar flat dark>
            <v-toolbar-title>New column</v-toolbar-title>
        </v-toolbar>
        <v-tabs vertical>
            <v-tab>
                Literal
            </v-tab>
            <v-tab-item>
                <vue-kodb-schema-literal-column
                    :libraryName="libraryName"

                    :confirm="confirm"
                >
                </vue-kodb-schema-literal-column>
            </v-tab-item>

            <v-tab>
                Reference
            </v-tab>
            <v-tab-item>
                <vue-kodb-schema-ref-column
                    :schema="schema"

                    :confirm="confirm"
                >
                </vue-kodb-schema-ref-column>
            </v-tab-item>

            <v-tab>
                List
            </v-tab>
            <v-tab-item>
                <vue-kodb-schema-list-column
                    :schema="schema"

                    :confirm="confirm"
                >
                </vue-kodb-schema-list-column>
            </v-tab-item>
        </v-tabs>
    </v-card>
</v-dialog>
`
})