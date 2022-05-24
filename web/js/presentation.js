import { network } from "./network.js"
export function createUserComponent(
    id,
    name,
    printPermission,
    pagesPerMonthPermission,
    userEmail,
) {
    // create fields
    const canPrintField = createCanPrintField(printPermission)
    const pagesPerMonthField = createPagesPerMonthField(pagesPerMonthPermission)
    const emailField = createEmailField(userEmail)
    
    // create controls
    const allowPrintingButton = createAllowPrintingButton();
    const forbidPrintingButton = createForbidPrintingButton();
    const setNumberOfPagesInput = createNumberOfPagesPerMonthInput();

    // add event handlers
    allowPrintingButton.addEventListener("click", async () => {
        canPrintField.innerText = await network.allowUsingPrinter(id);
    });

    forbidPrintingButton.addEventListener("click", async () => {
        canPrintField.innerText = (await network.forbidUsingPrinter(id));
    });

    // add user form
    setNumberOfPagesInput.addEventListener("submit", async (e) => {
        e.preventDefault();
        pagesPerMonthField.innerText = await network.setNumberOfPages(id, Number(setNumberOfPagesInput.firstChild.value));
        setNumberOfPagesInput.firstChild.value = "";
    });

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

    // add body
    {
        const userManagement = document.createElement("div");
        userManagement.classList.add("user-managment");

        // add user details
        const userDetails = createUserDetails(emailField, canPrintField, pagesPerMonthField);
        userManagement.appendChild(userDetails);

        // add user controls
        const userControls = createUserControls(allowPrintingButton, forbidPrintingButton, setNumberOfPagesInput);
        userManagement.appendChild(userControls);

        user.appendChild(userManagement);
    }

    return user;
}

function createUserDetails(emailField, printPermissionField, pagesPerMonthField) {
    const userDetails = document.createElement("div");
    userDetails.classList.add("user-details");

    userDetails.appendChild(printPermissionField)
    userDetails.appendChild(pagesPerMonthField)
    userDetails.appendChild(emailField)

    return userDetails 
}

function createEmailField(userEmail) {
    const email = document.createElement("p");
    email.innerText = userEmail;
    email.classList.add("user-info");

    return email
}

function createCanPrintField(printPermission) {
    const canPrintField = document.createElement("p");
    canPrintField.innerText = printPermission;
    canPrintField.classList.add("user-info");

    return  canPrintField;
}

function createPagesPerMonthField(pagesPerMonthPermission) {
    const pagesPerMonthField = document.createElement("p");
    pagesPerMonthField.innerText = pagesPerMonthPermission;
    pagesPerMonthField.classList.add("user-info");

    return pagesPerMonthField
}

function createUserControls(
    allowPrintingButton,
    forbidPrintingButton,
    setNumberOfPagesInput
) {
    const userControls = document.createElement("div");
    userControls.classList.add("user-controls");

    // add userButtons
    {
        const userButtons = document.createElement("div");
        userButtons.classList.add("user-buttons");

        userButtons.appendChild(allowPrintingButton);
        userButtons.appendChild(forbidPrintingButton);

        userControls.appendChild(userButtons);
    }

    userControls.appendChild(setNumberOfPagesInput);

    return userControls;
}

function createAllowPrintingButton() {
    const allowPrintingButton = document.createElement("button");

    allowPrintingButton.classList.add("user-button");
    allowPrintingButton.classList.add("user-allow-printing");
    allowPrintingButton.innerText = "Allow Printing";

    return allowPrintingButton;
}

function createForbidPrintingButton() {
    const forbidPrintingButton = document.createElement("button");

    forbidPrintingButton.classList.add("user-button");
    forbidPrintingButton.classList.add("user-forbid-printing");
    forbidPrintingButton.innerText = "Forbid Printing"

    return forbidPrintingButton;
}

function createNumberOfPagesPerMonthInput() {
    const pageInput = document.createElement("form");
    pageInput.classList.add("user-form");

    {
        const pageInputBox = document.createElement("input");
        pageInputBox.type = "number";
        pageInputBox.placeholder = "Set Number Of Pages Per Month";
        pageInput.appendChild(pageInputBox);

        const submit = document.createElement("input");
        submit.type = "submit";
        pageInput.appendChild(submit);
    }

    return pageInput;
}