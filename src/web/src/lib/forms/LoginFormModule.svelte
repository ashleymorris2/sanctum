<script lang="ts" context="module">
    interface ServerResult {
        type: 'success' | 'failure';
        data: {
            email?: string;
            password?: string;
            error?: string;
            message?: string;
            authToken?: string;
            location?: string;
        };
    }

    export interface FormErrors {
        email?: string;
        password?: string;
        error?: string;
        message?: string;
    }

    export function createLoginForm() {
        const formData = $state({
            email: '',
            password: '',
            errors: {} as FormErrors
        });

        const clearErrors = (field?: keyof FormErrors) => {
            if (field) {
                delete formData.errors[field];
                delete formData.errors.message;
            } else {
                formData.errors = {};
            }
        };

        const setError = (field: keyof FormErrors, message: string) => {
            formData.errors[field] = message;
        };

        const validateForm = (cancel?: () => void) => {
            clearErrors();

            const emailValidation = validateEmail(formData.email);
            if (!emailValidation) {
                if (cancel) cancel();
                return false;
            }

            if (!formData.password) {
                if (cancel) cancel();
                setError('password', 'Password is required');
                return false;
            }

            return true;
        };

        const handleServerErrors = (result: ServerResult) => {
            if (result.data.email) setError('email', result.data.email);
            if (result.data.password) setError('password', result.data.password);
            if (result.data.error) setError('error', result.data.error);
            if (result.data.message) setError('message', result.data.message);
        };

        const validateEmail = (email: string): boolean => {
            if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
                setError('email', 'Please enter a valid email');
                return false;
            }
            clearErrors('email');
            clearErrors('message');
            return true;
        };

        return {
            formData,
            validateForm,
            clearErrors,
            handleServerErrors
        };
    }
</script>