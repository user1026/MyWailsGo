// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {cpu} from '../models';
import {Computer} from '../models';

export function GetAllInfo():Promise<void>;

export function GetCpuInfo():Promise<cpu.InfoStat>;

export function GetMemInfo():Promise<Computer.MemInfo>;

export function GetUsingCpuInfo():Promise<number>;

export function Greet(arg1:string):Promise<string>;
