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
                NumberOfPages: numberOfPages,
            }),
        });

        return JSON.parse(JSON.stringify(await response.json()));
    },

    allowUsingPrinter: async function (userID) {
        const response = await fetch(`/users/${userID}/printing/permission`, {
            method: "PATCH",
        });

        return JSON.parse(JSON.stringify(await response.json()));
    },

    forbidUsingPrinter: async function (userID) {
        const response = await fetch(`/users/${userID}/printing/prohibition`, {
            method: "PATCH",
        });

        return JSON.parse(JSON.stringify(await response.json()));
    },
}