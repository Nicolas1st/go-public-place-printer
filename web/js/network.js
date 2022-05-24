export const network = {
    getAllUsers: async function () {
        const response = await fetch("/users/");
        return JSON.parse(JSON.stringify(await response.json()));
    },

    setNumberOfPages: async function (userID, numberOfPages) {
        const response = await fetch(`/users/${userID}/pages`, {
            method: "PATCH",
            header: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                numberOfPages: numberOfPages,
            }),
        });

        const resp = JSON.parse(JSON.stringify(await response.json()));

        return resp.numberOfPages;
    },

    allowUsingPrinter: async function (userID) {
        const response = await fetch(`/users/${userID}/printing/permission`, {
            method: "PATCH",
        });

        const resp = JSON.parse(JSON.stringify(await response.json()));

        return resp.permission;
    },

    forbidUsingPrinter: async function (userID) {
        const response = await fetch(`/users/${userID}/printing/prohibition`, {
            method: "PATCH",
        });

        const resp = JSON.parse(JSON.stringify(await response.json()));

        return resp.permission;
    }
}