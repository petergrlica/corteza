<template>
  <wrap
    v-bind="$props"
    v-on="$listeners"
  >
    <div
      v-if="!options.tabs.length"
      class="d-flex h-100 align-items-center justify-content-center"
    >
      <p class="text-secondary mb-0">
        {{ $t('tabs.noTabsBase') }}
      </p>
    </div>
    <b-tabs
      v-else
      v-model="activeTab"
      v-bind="{
        align: options.style.alignment,
        fill: options.style.fillJustify === 'fill',
        justified: options.style.fillJustify === 'justified',
        pills: options.style.appearance === 'pills',
        tabs: options.style.appearance === 'tabs',
        small: options.style.appearance === 'small',
        vertical: options.style.verticalHorizontal === 'vertical',
      }"
    >
      <b-tab
        v-for="(tab, index) in options.tabs"
        :key="index"
        :title="tab.block.title"
        @click="trackTab(index)"
      >
        <page-block-tab
          v-if="index === activeTab"
          v-bind="{ ...$attrs, ...$props, page, block: tab.block, blockIndex: index }"
          :record="record"
          :module="module"
        />
      </b-tab>
    </b-tabs>
  </wrap>
</template>

<script>
import base from './base'

export default {
  i18nOptions: {
    namespaces: 'block',
  },

  name: 'TabBase',

  components: {
    PageBlockTab: () => import('corteza-webapp-compose/src/components/PageBlocks'),
  },
  extends: base,

  data () {
    return {
      activeTab: 0,
    }
  },

  watch: {
    '$route.query': {
      immediate: true,

      handler (query) {
        this.changeActiveTab(query)
      },
    },
  },

  methods: {
    trackTab (index) {
      this.$router.replace({ query: { tabBlockIndex: this.options.blockIndex, tabIndex: index } })
    },

    changeActiveTab (query) {
      if (+query.tabBlockIndex === this.options.blockIndex) {
        this.activeTab = +query.tabIndex
      }
    },
  },
}
</script>
