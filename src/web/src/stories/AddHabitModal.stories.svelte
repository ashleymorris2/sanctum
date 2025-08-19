<script module lang="ts">
    import {defineMeta} from '@storybook/addon-svelte-csf';
    import AddHabitModal from "../components/AddHabitModal.svelte";

    const {Story} = defineMeta({
        title: 'Modals/AddHabitModal',
        component: AddHabitModal,
        parameters: {
            layout: 'fullscreen',
        },
        args: {
            open: {control: 'boolean'}
        },
    });
</script>

<Story name="Default" let:args>
    <AddHabitModal {...args}/>
</Story>