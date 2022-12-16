<template>
  <div>
    <page-block
      v-bind="{ ...$attrs, ...$props, page, block, boundingRect, blockIndex: index }"
      :record="record"
      :module="module"
    />
  </div>
</template>

<script>
import { compose } from '@cortezaproject/corteza-js'
export default {

  components: {
    PageBlock: () => import('corteza-webapp-compose/src/components/PageBlocks'),
  },
  props: {
    boundingRect: {
      type: Object,
      required: false,
      default: undefined,
    },

    blockIndex: {
      type: Number,
      default: -1,
    },

    index: {
      type: Number,
      default: -1,
    },

    namespace: {
      type: compose.Namespace,
      required: true,
    },

    page: {
      type: compose.Page,
      required: true,
    },

    block: {
      type: compose.PageBlock,
      required: true,
    },

    module: {
      type: compose.Module,
      required: false,
      default: undefined,
    },

    record: {
      type: compose.Record,
      required: false,
      default: undefined,
    },

    mode: {
      type: String,
      required: false,
      default: '',
    },
  },

}
</script>
