Vue.component("kodb-library", {
        data: function() {
                return {
                        headers:[],
                        desserts:[],
                }
        },
        template:
`
<v-data-table
        :headers="headers"
        :items="desserts"
        :items-per-page="10"
></v-data-table>
`
});