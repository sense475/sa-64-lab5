import { UsersInterface } from "./IUser";
import { JobInterface } from "./IJob";
import { SexInterface } from "./ISex";
import { InsuranceInterface } from "./IIns";

export interface PatientInterface {

    ID: number,
    Firstname: string,
    Lastname: string,
    Age: number,
    IDcard: string,
    Time: Date,
    Tel: string,
    UserNurseID: number,
	UserNurse:   UsersInterface,

	JobID: number,
	Job:   JobInterface,

	InsuranceID: number,
	Insurance:   InsuranceInterface,

	SexID: number,
	Sex:   SexInterface,

   }