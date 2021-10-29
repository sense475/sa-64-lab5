import { RolesInterface } from "./IRole";
export interface UsersInterface {

    ID: number,
    Name: string,
    Username: string,
    Password: String,
    RoleID: number,
    Role: RolesInterface,
}