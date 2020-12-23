<template>
  <NotificationButton
    color="blue"
    :placeholder="'Payer'"
    :title="'Notification'"
    :url="'/p4'"
    :start="3"
    :load="true"
  ></NotificationButton>
</template>


<script>
import constants from "@/constants";
import NotificationButton from "@/components/NotificationButton";
export default {
  components: {
    NotificationButton
  }
};
</script>