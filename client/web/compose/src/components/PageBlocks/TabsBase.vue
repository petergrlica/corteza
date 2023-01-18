<template>
  <wrap
    v-bind="$props"
    :scrollable-body="false"
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
      content-class="flex-fill"
      v-bind="{
        align: block.options.style.alignment,
        fill: block.options.style.fillJustify === 'fill',
        justified: block.options.style.fillJustify === 'justified',
        pills: block.options.style.appearance === 'pills',
        tabs: block.options.style.appearance === 'tabs',
        small: block.options.style.appearance === 'small',
        vertical: block.options.style.verticalHorizontal === 'vertical',
      }"
      class="ml-2 mt-3 mr-2 d-flex flex-column h-100"
    >
      <b-tab
        v-for="(tab, index) in options.tabs"
        :key="index"
        :title="tab.block.title ? tab.block.title : `Block-${tab.block.kind}-${tab.indexOnMain}`"
        class="h-100"
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
