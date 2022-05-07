export function createUserComponent(
    id,
    name,
    printPermission,
    pagesPerMonthPermission,
    userEmail,
    allowFunction,
    forbidFunction,
    setPagesFunction
) {
    // create user element
    const user = document.createElement("div");
    user.classList.add("user");

    // add name
    {
        const userName = document.createElement("p");
        userName.classList.add("user-name");
        userName.innerText = name;
        user.appendChild(userName)
    }

    // add user managment
    {
        const userManagment = document.createElement("div");
        userManagment.classList.add("user-managment");
        user.appendChild(userManagment);

        // add user info
        const userDetails = createUserDetails(printPermission, pagesPerMonthPermission, userEmail);
        userManagment.appendChild(userDetails);

        // add user controls
        const userControls = createUserControls(id, allowFunction, forbidFunction, setPagesFunction);
        userManagment.appendChild(userControls);
    }

    return user;
}

function createUserDetails(printPermission, pagesPerMonthPermission, userEmail) {
    const userDetails = document.createElement("div");
    userDetails.classList.add("user-details");

    const canPrint = document.createElement("p");
    canPrint.innerText = printPermission;
    canPrint.classList.add("user-info");
    userDetails.appendChild(canPrint);

    const pagesPerMonth = document.createElement("p");
    pagesPerMonth.innerText = pagesPerMonthPermission;
    pagesPerMonth.classList.add("user-info");
    userDetails.appendChild(pagesPerMonth);

    const Email = document.createElement("p");
    Email.innerText = userEmail;
    Email.classList.add("user-info");
    userDetails.appendChild(Email);

    return userDetails;
}

function createUserControls(id, allowFunction, forbidFunction, setPagesFunction) {
    const userControls = document.createElement("div");
    userControls.classList.add("user-controls");

    // add userButtons
    const userButtons = document.createElement("div");
    userButtons.classList.add("user-buttons");
    userControls.appendChild(userButtons);

    // add allow printing button
    const allowPrintingButton = document.createElement("button");
    allowPrintingButton.classList.add("user-button");
    allowPrintingButton.classList.add("user-allow-printing");
    allowPrintingButton.innerText = "Allow Printing"
    allowPrintingButton.addEventListener("click", () => {
        allowFunction(id);
    });
    userButtons.appendChild(allowPrintingButton);

    // add forbid printing button
    const forbidPrintingButton = document.createElement("button");
    forbidPrintingButton.classList.add("user-button");
    forbidPrintingButton.classList.add("user-forbid-printing");
    forbidPrintingButton.innerText = "Forbid Printing"
    forbidPrintingButton.addEventListener("click", () => {
        forbidFunction(id);
    });
    userButtons.appendChild(forbidPrintingButton);

    // add user form
    const userForm = document.createElement("form");
    userForm.classList.add("user-form");
    userControls.appendChild(userForm);
    
    const pageInput = document.createElement("input");
    pageInput.type = "number";
    pageInput.name = "nameNumberOfPages";
    pageInput.placeholder = "Set Number Of Pages Per Month";
    userForm.appendChild(pageInput);

    userForm.addEventListener("submit", (e) => {
        e.preventDefault();
        setPagesFunction(id, Number(pageInput.value));
        pageInput.value = "";
    });

    return userControls;
}
