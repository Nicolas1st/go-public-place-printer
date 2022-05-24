import { createUserComponent } from "./presentation.js";
import { network } from "./network";

const  userManager = document.querySelector(".user-manager");
document.addEventListener("DOMContentLoaded", async () => {
    const users = await network.getAllUsers();
    users.forEach(data => {
        const user = createUserComponent(
            data.ID,
            data.Name,
            data.CanUsePrinter,
            data.PagesPerMonth,
            data.Email,
        );

        userManager.appendChild(user);
    });
});
