import { UsersInterface } from "./IUser";
import { PatientInterface } from "./IPatient";
import { RemedyInterface } from "../models/IRemedy";
export interface PaymentInterface {
    ID: number,
    Price: number,
    Note: String,
    UserFinancialID: number,
    UserFinancial: UsersInterface,
    PatientID: number,
    Patient: PatientInterface,
    RemedyTypeID: number,
    RemedyType: RemedyInterface,

    Paytime: Date,
}