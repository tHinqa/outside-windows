package wincrypt

import . "github.com/tHinqa/outside"
import . "github.com/tHinqa/outside-windows/types"

var CryptAcquireContext func(
	prov *HCRYPTPROV,
	container, provider VString,
	provType, flags DWORD) BOOL

var CryptReleaseContext func(prov HCRYPTPROV, flags DWORD) BOOL

var CryptGenKey func(
	Prov HCRYPTPROV,
	Algid ALG_ID,
	Flags DWORD,
	Key *HCRYPTKEY) BOOL

var CryptDeriveKey func(
	Prov HCRYPTPROV,
	Algid ALG_ID,
	BaseData HCRYPTHASH,
	Flags DWORD,
	Key *HCRYPTKEY) BOOL

var CryptDestroyKey func(
	Key HCRYPTKEY) BOOL

var CryptSetKeyParam func(
	Key HCRYPTKEY,
	Param DWORD,
	Data *BYTE,
	Flags DWORD) BOOL

var CryptGetKeyParam func(
	Key HCRYPTKEY,
	Param DWORD,
	Data *BYTE,
	DataLen *DWORD,
	Flags DWORD) BOOL

var CryptSetHashParam func(
	hHash HCRYPTHASH,
	dwParam DWORD,
	pbData *BYTE,
	dwFlags DWORD) BOOL

var CryptGetHashParam func(
	hHash HCRYPTHASH,
	dwParam DWORD,
	pbData *BYTE,
	pdwDataLen *DWORD,
	dwFlags DWORD) BOOL

var CryptSetProvParam func(
	hProv HCRYPTPROV,
	dwParam DWORD,
	pbData *BYTE,
	dwFlags DWORD) BOOL

var CryptGetProvParam func(
	hProv HCRYPTPROV,
	dwParam DWORD,
	pbData *BYTE,
	pdwDataLen *DWORD,
	dwFlags DWORD) BOOL

var CryptGenRandom func(
	hProv HCRYPTPROV,
	dwLen DWORD,
	pbBuffer *BYTE) BOOL

var CryptGetUserKey func(
	hProv HCRYPTPROV,
	dwKeySpec DWORD,
	phUserKey *HCRYPTKEY) BOOL

var CryptExportKey func(
	hKey HCRYPTKEY,
	hExpKey HCRYPTKEY,
	dwBlobType DWORD,
	dwFlags DWORD,
	pbData *BYTE,
	pdwDataLen *DWORD) BOOL

var CryptImportKey func(
	hProv HCRYPTPROV,
	pbData *BYTE,
	dwDataLen DWORD,
	hPubKey HCRYPTKEY,
	dwFlags DWORD,
	phKey *HCRYPTKEY) BOOL

var CryptEncrypt func(
	hKey HCRYPTKEY,
	hHash HCRYPTHASH,
	Final BOOL,
	dwFlags DWORD,
	pbData *BYTE,
	pdwDataLen *DWORD,
	dwBufLen DWORD) BOOL

var CryptDecrypt func(
	hKey HCRYPTKEY,
	hHash HCRYPTHASH,
	Final BOOL,
	dwFlags DWORD,
	pbData *BYTE,
	pdwDataLen *DWORD) BOOL

var CryptCreateHash func(
	hProv HCRYPTPROV,
	Algid ALG_ID,
	hKey HCRYPTKEY,
	dwFlags DWORD,
	phHash *HCRYPTHASH) BOOL

var CryptHashData func(
	hHash HCRYPTHASH,
	pbData *BYTE,
	dwDataLen DWORD,
	dwFlags DWORD) BOOL

var CryptHashSessionKey func(
	hHash HCRYPTHASH,
	hKey HCRYPTKEY,
	dwFlags DWORD) BOOL

var CryptDestroyHash func(
	hHash HCRYPTHASH) BOOL

var CryptSignHash func(
	hash HCRYPTHASH,
	keySpec DWORD,
	description VString,
	flags DWORD,
	signature *BYTE,
	sigLen *DWORD) BOOL

var CryptVerifySignature func(
	hash HCRYPTHASH,
	signature *BYTE,
	sigLen DWORD,
	pubKey HCRYPTKEY,
	description VString,
	flags DWORD) BOOL

var CryptSetProvider func(provName VString, provType DWORD) BOOL

var CryptSetProviderEx func(
	provName VString,
	provType DWORD,
	reserved *DWORD,
	flags DWORD) BOOL

var CryptGetDefaultProvider func(
	provType DWORD,
	reserved *DWORD,
	flags DWORD,
	provName OVString,
	provNameSize *DWORD) BOOL

var CryptEnumProviderTypes func(
	index DWORD,
	reserved *DWORD,
	flags DWORD,
	provType *DWORD,
	typeName OVString,
	typeNameSize *DWORD) BOOL

var CryptEnumProviders func(
	index DWORD,
	reserved *DWORD,
	flags DWORD,
	provType *DWORD,
	provName OVString,
	provNameSize *DWORD) BOOL

var CryptContextAddRef func(
	prov HCRYPTPROV, reserved *DWORD, flags DWORD) BOOL

var CryptDuplicateKey func(
	key HCRYPTKEY,
	reserved *DWORD,
	flags DWORD,
	newKey *HCRYPTKEY) BOOL

var CryptDuplicateHash func(
	hash HCRYPTHASH,
	reserved *DWORD,
	flags DWORD,
	newHash *HCRYPTHASH) BOOL

//WINCRYPT32API
//BOOL
//WINAPI
//CryptEncodeObjectEx(
//    IN DWORD dwCertEncodingType,
//    IN LPCSTR lpszStructType,
//    IN const void *pvStructInfo,
//    IN DWORD dwFlags,
//    IN OPTIONAL PCRYPT_ENCODE_PARA pEncodePara,
//    OUT void *pvEncoded,
//    IN OUT DWORD *pcbEncoded
//    );

//WINCRYPT32API
//BOOL
//WINAPI
//CryptEncodeObject(
//    IN DWORD        dwCertEncodingType,
//    IN LPCSTR       lpszStructType,
//    IN const void   *pvStructInfo,
//    OUT BYTE        *pbEncoded,
//    IN OUT DWORD    *pcbEncoded
//    );

//WINCRYPT32API
//BOOL
//WINAPI
//CryptDecodeObjectEx(
//    IN DWORD dwCertEncodingType,
//    IN LPCSTR lpszStructType,
//    IN const BYTE *pbEncoded,
//    IN DWORD cbEncoded,
//    IN DWORD dwFlags,
//    IN OPTIONAL PCRYPT_DECODE_PARA pDecodePara,
//    OUT OPTIONAL void *pvStructInfo,
//    IN OUT DWORD *pcbStructInfo
//    );

//WINCRYPT32API
//BOOL
//WINAPI
//CryptDecodeObject(
//    IN DWORD        dwCertEncodingType,
//    IN LPCSTR       lpszStructType,
//    IN const BYTE   *pbEncoded,
//    IN DWORD        cbEncoded,
//    IN DWORD        dwFlags,
//    OUT void        *pvStructInfo,
//    IN OUT DWORD    *pcbStructInfo
//    );

//WINCRYPT32API
//BOOL
//WINAPI
//CryptInstallOIDFunctionAddress(
//    IN HMODULE hModule,         // hModule passed to DllMain
//    IN DWORD dwEncodingType,
//    IN LPCSTR pszFuncName,
//    IN DWORD cFuncEntry,
//    IN const CRYPT_OID_FUNC_ENTRY rgFuncEntry[],
//    IN DWORD dwFlags
//    );

//WINCRYPT32API
//HCRYPTOIDFUNCSET
//WINAPI
//CryptInitOIDFunctionSet(
//    IN LPCSTR pszFuncName,
//    IN DWORD dwFlags
//    );

//WINCRYPT32API
//BOOL
//WINAPI
//CryptGetOIDFunctionAddress(
//    IN HCRYPTOIDFUNCSET hFuncSet,
//    IN DWORD dwEncodingType,
//    IN LPCSTR pszOID,
//    IN DWORD dwFlags,
//    OUT void **ppvFuncAddr,
//    OUT HCRYPTOIDFUNCADDR *phFuncAddr
//    );

//WINCRYPT32API
//BOOL
//WINAPI
//CryptGetDefaultOIDDllList(
//    IN HCRYPTOIDFUNCSET hFuncSet,
//    IN DWORD dwEncodingType,
//    OUT LPWSTR pwszDllList,
//    IN OUT DWORD *pcchDllList
//    );

//WINCRYPT32API
//BOOL
//WINAPI
//CryptGetDefaultOIDFunctionAddress(
//    IN HCRYPTOIDFUNCSET hFuncSet,
//    IN DWORD dwEncodingType,
//    IN OPTIONAL LPCWSTR pwszDll,
//    IN DWORD dwFlags,
//    OUT void **ppvFuncAddr,
//    IN OUT HCRYPTOIDFUNCADDR *phFuncAddr
//    );

//WINCRYPT32API
//BOOL
//WINAPI
//CryptFreeOIDFunctionAddress(
//    IN HCRYPTOIDFUNCADDR hFuncAddr,
//    IN DWORD dwFlags
//    );

//WINCRYPT32API
//BOOL
//WINAPI
//CryptRegisterOIDFunction(
//    IN DWORD dwEncodingType,
//    IN LPCSTR pszFuncName,
//    IN LPCSTR pszOID,
//    IN OPTIONAL LPCWSTR pwszDll,
//    IN OPTIONAL LPCSTR pszOverrideFuncName
//    );

//WINCRYPT32API
//BOOL
//WINAPI
//CryptUnregisterOIDFunction(
//    IN DWORD dwEncodingType,
//    IN LPCSTR pszFuncName,
//    IN LPCSTR pszOID
//    );

//WINCRYPT32API
//BOOL
//WINAPI
//CryptRegisterDefaultOIDFunction(
//    IN DWORD dwEncodingType,
//    IN LPCSTR pszFuncName,
//    IN DWORD dwIndex,
//    IN LPCWSTR pwszDll
//    );

//WINCRYPT32API
//BOOL
//WINAPI
//CryptUnregisterDefaultOIDFunction(
//    IN DWORD dwEncodingType,
//    IN LPCSTR pszFuncName,
//    IN LPCWSTR pwszDll
//    );

//WINCRYPT32API
//BOOL
//WINAPI
//CryptSetOIDFunctionValue(
//    IN DWORD dwEncodingType,
//    IN LPCSTR pszFuncName,
//    IN LPCSTR pszOID,
//    IN LPCWSTR pwszValueName,
//    IN DWORD dwValueType,
//    IN const BYTE *pbValueData,
//    IN DWORD cbValueData
//    );

//WINCRYPT32API
//BOOL
//WINAPI
//CryptGetOIDFunctionValue(
//    IN DWORD dwEncodingType,
//    IN LPCSTR pszFuncName,
//    IN LPCSTR pszOID,
//    IN LPCWSTR pwszValueName,
//    OUT DWORD *pdwValueType,
//    OUT BYTE *pbValueData,
//    IN OUT DWORD *pcbValueData
//    );

//WINCRYPT32API
//BOOL
//WINAPI
//CryptFormatObject(
//    IN DWORD dwCertEncodingType,
//    IN DWORD dwFormatType,
//    IN DWORD dwFormatStrType,
//    IN void  *pFormatStruct,
//    IN LPCSTR lpszStructType,
//    IN const BYTE *pbEncoded,
//    IN DWORD cbEncoded,
//    OUT void *pbFormat,
//    IN OUT DWORD *pcbFormat
//    );
