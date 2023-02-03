import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class HashService {

  constructor() { }

  public static async getHash(password: string){
    var hash = await this.digest(password);
    var hashString = this.digestToString(hash);
    return hashString;
  }

  private static async digest(password: string){
    var encoder = new TextEncoder();
    var encoded = encoder.encode(password.normalize())
    var hash = await crypto.subtle.digest("SHA-256", encoded);
    return hash
  }

  private static digestToString(buffer: ArrayBuffer){
    const byteArray = new Uint8Array(buffer);

    const hexCodes = [...byteArray].map(value => {
      const hexCode = value.toString(16);
      const paddedHexCode = hexCode.padStart(2, '0');
      return paddedHexCode;
    });
  
    return hexCodes.join('');
  }
}
