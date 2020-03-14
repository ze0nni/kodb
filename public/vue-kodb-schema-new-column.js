Vue.component("vue-kodb-schema-new-column", {
        props: [
            "schema"
        ],
        template:
    `
<v-dialog>
    <template v-slot:activator="{ on }">
        <v-btn text block  v-on="on">
            <v-icon left>mdi-plus</v-icon>
            New
        </v-btn>
    </template>

    <v-card>
        <v-toolbar flat dark>
            <v-toolbar-title>New column</v-toolbar-title>
        </v-toolbar>
        <v-tabs vertical>
            <v-tab>
                Literal
            </v-tab>
            <v-tab-item>
                <v-card>
                <v-col>
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
                    </v-col>
                </v-card>
            </v-tab-item>

            <v-tab>
                Reference
            </v-tab>
            <v-tab-item>
                <v-card>
                    <v-col>
                        <v-text-field
                        >
                        </v-text-field>

                        <v-select
                            :items="schema"
                        >
                        </v-select>
                    </v-col>
                </v-card>
            </v-tab-item>

            <v-tab>
                List
            </v-tab>
            <v-tab-item>
                <v-card>
                    <v-text-field
                    >
                    </v-text-field>

                    <v-col>
                        <v-select
                            :items="schema"
                        >
                        </v-select>
                    </v-col>
                </v-card>
            </v-tab-item>
        </v-tabs>
    </v-card>
</v-dialog>
`
})