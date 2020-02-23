Vue.component("kodb", {
        template:
`
<v-app>
        <v-content>
                <v-tabs
                        show-arrows
                >
                        <v-tab
                                v-for="i in 5"
                                :key="i"
                                :href="'#tab-' + i"
                        >
                                Item {{ i }}
                        </v-tab>
                </v-tabs>
                <v-toolbar>
                        <v-btn text>+</v-btn>
                        <v-btn text>-</v-btn>
                </v-toolbar>
                <v-container>Hello world</v-container>
        </v-content>
</v-app>
`
})