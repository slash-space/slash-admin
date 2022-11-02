// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package generated

import (
	converter "slash-admin/app/admin/cmd/api/internal/converter"
	types "slash-admin/app/admin/cmd/api/internal/types"
	ent "slash-admin/app/admin/ent"
)

type ConverterImpl struct{}

func (c *ConverterImpl) ConvertPagination(source *ent.PageDetails) *types.Pagination {
	var pTypesPagination *types.Pagination
	if source != nil {
		typesPagination := c.entPageDetailsToTypesPagination(*source)
		pTypesPagination = &typesPagination
	}
	return pTypesPagination
}
func (c *ConverterImpl) ConvertSysApi(source *ent.SysApi) *types.ApiInfo {
	var pTypesApiInfo *types.ApiInfo
	if source != nil {
		typesApiInfo := c.entSysApiToTypesApiInfo(*source)
		pTypesApiInfo = &typesApiInfo
	}
	return pTypesApiInfo
}
func (c *ConverterImpl) ConvertSysApiList(source []*ent.SysApi) []*types.ApiInfo {
	pTypesApiInfoList := make([]*types.ApiInfo, len(source))
	for i := 0; i < len(source); i++ {
		pTypesApiInfoList[i] = c.ConvertSysApi(source[i])
	}
	return pTypesApiInfoList
}
func (c *ConverterImpl) ConvertSysDictionary(source *ent.SysDictionary) *types.DictionaryInfo {
	var pTypesDictionaryInfo *types.DictionaryInfo
	if source != nil {
		typesDictionaryInfo := c.entSysDictionaryToTypesDictionaryInfo(*source)
		pTypesDictionaryInfo = &typesDictionaryInfo
	}
	return pTypesDictionaryInfo
}
func (c *ConverterImpl) ConvertSysDictionaryDetail(source *ent.SysDictionaryDetail) *types.DictionaryDetailInfo {
	var pTypesDictionaryDetailInfo *types.DictionaryDetailInfo
	if source != nil {
		typesDictionaryDetailInfo := c.entSysDictionaryDetailToTypesDictionaryDetailInfo(*source)
		pTypesDictionaryDetailInfo = &typesDictionaryDetailInfo
	}
	return pTypesDictionaryDetailInfo
}
func (c *ConverterImpl) ConvertSysDictionaryDetailList(source []*ent.SysDictionaryDetail) []*types.DictionaryDetailInfo {
	pTypesDictionaryDetailInfoList := make([]*types.DictionaryDetailInfo, len(source))
	for i := 0; i < len(source); i++ {
		pTypesDictionaryDetailInfoList[i] = c.ConvertSysDictionaryDetail(source[i])
	}
	return pTypesDictionaryDetailInfoList
}
func (c *ConverterImpl) ConvertSysDictionaryList(source []*ent.SysDictionary) []*types.DictionaryInfo {
	pTypesDictionaryInfoList := make([]*types.DictionaryInfo, len(source))
	for i := 0; i < len(source); i++ {
		pTypesDictionaryInfoList[i] = c.ConvertSysDictionary(source[i])
	}
	return pTypesDictionaryInfoList
}
func (c *ConverterImpl) ConvertSysOauthProvider(source *ent.SysOauthProvider) *types.OauthProviderInfo {
	var pTypesOauthProviderInfo *types.OauthProviderInfo
	if source != nil {
		typesOauthProviderInfo := c.entSysOauthProviderToTypesOauthProviderInfo(*source)
		pTypesOauthProviderInfo = &typesOauthProviderInfo
	}
	return pTypesOauthProviderInfo
}
func (c *ConverterImpl) ConvertSysOauthProviderList(source []*ent.SysOauthProvider) []*types.OauthProviderInfo {
	pTypesOauthProviderInfoList := make([]*types.OauthProviderInfo, len(source))
	for i := 0; i < len(source); i++ {
		pTypesOauthProviderInfoList[i] = c.ConvertSysOauthProvider(source[i])
	}
	return pTypesOauthProviderInfoList
}
func (c *ConverterImpl) ConvertSysRoleToRoleInfo(source *ent.SysRole) *types.RoleInfo {
	var pTypesRoleInfo *types.RoleInfo
	if source != nil {
		typesRoleInfo := c.entSysRoleToTypesRoleInfo(*source)
		pTypesRoleInfo = &typesRoleInfo
	}
	return pTypesRoleInfo
}
func (c *ConverterImpl) ConvertSysRoleToRoleInfoList(source []*ent.SysRole) []*types.RoleInfo {
	pTypesRoleInfoList := make([]*types.RoleInfo, len(source))
	for i := 0; i < len(source); i++ {
		pTypesRoleInfoList[i] = c.ConvertSysRoleToRoleInfo(source[i])
	}
	return pTypesRoleInfoList
}
func (c *ConverterImpl) ConvertSysToken(source *ent.SysToken) *types.TokenInfo {
	var pTypesTokenInfo *types.TokenInfo
	if source != nil {
		typesTokenInfo := c.entSysTokenToTypesTokenInfo(*source)
		pTypesTokenInfo = &typesTokenInfo
	}
	return pTypesTokenInfo
}
func (c *ConverterImpl) ConvertSysTokenList(source []*ent.SysToken) []*types.TokenInfo {
	pTypesTokenInfoList := make([]*types.TokenInfo, len(source))
	for i := 0; i < len(source); i++ {
		pTypesTokenInfoList[i] = c.ConvertSysToken(source[i])
	}
	return pTypesTokenInfoList
}
func (c *ConverterImpl) ConvertSysUser(source *ent.SysUser) *types.UserInfo {
	var pTypesUserInfo *types.UserInfo
	if source != nil {
		typesUserInfo := c.entSysUserToTypesUserInfo(*source)
		pTypesUserInfo = &typesUserInfo
	}
	return pTypesUserInfo
}
func (c *ConverterImpl) ConvertSysUserList(source []*ent.SysUser) []*types.UserInfo {
	pTypesUserInfoList := make([]*types.UserInfo, len(source))
	for i := 0; i < len(source); i++ {
		pTypesUserInfoList[i] = c.ConvertSysUser(source[i])
	}
	return pTypesUserInfoList
}
func (c *ConverterImpl) entPageDetailsToTypesPagination(source ent.PageDetails) types.Pagination {
	var typesPagination types.Pagination
	typesPagination.Page = source.Page
	typesPagination.Limit = source.Limit
	typesPagination.Total = source.Total
	return typesPagination
}
func (c *ConverterImpl) entSysApiToTypesApiInfo(source ent.SysApi) types.ApiInfo {
	var typesApiInfo types.ApiInfo
	typesApiInfo.ID = source.ID
	typesApiInfo.Path = source.Path
	typesApiInfo.Description = source.Description
	typesApiInfo.Group = source.APIGroup
	typesApiInfo.Method = source.Method
	typesApiInfo.CreatedAt = converter.TimeToUnixMilli(source.CreatedAt)
	return typesApiInfo
}
func (c *ConverterImpl) entSysDictionaryDetailToTypesDictionaryDetailInfo(source ent.SysDictionaryDetail) types.DictionaryDetailInfo {
	var typesDictionaryDetailInfo types.DictionaryDetailInfo
	typesDictionaryDetailInfo.ID = source.ID
	typesDictionaryDetailInfo.Title = source.Title
	typesDictionaryDetailInfo.Key = source.Key
	typesDictionaryDetailInfo.Value = source.Value
	typesDictionaryDetailInfo.Status = converter.StatusToUint8(source.Status)
	typesDictionaryDetailInfo.CreatedAt = converter.TimeToUnixMilli(source.CreatedAt)
	typesDictionaryDetailInfo.UpdatedAt = converter.TimeToUnixMilli(source.UpdatedAt)
	typesDictionaryDetailInfo.DictionaryID = source.DictionaryID
	return typesDictionaryDetailInfo
}
func (c *ConverterImpl) entSysDictionaryToTypesDictionaryInfo(source ent.SysDictionary) types.DictionaryInfo {
	var typesDictionaryInfo types.DictionaryInfo
	typesDictionaryInfo.ID = source.ID
	typesDictionaryInfo.Title = source.Title
	typesDictionaryInfo.Name = source.Name
	typesDictionaryInfo.Status = converter.StatusToUint8(source.Status)
	typesDictionaryInfo.Description = source.Desc
	typesDictionaryInfo.CreatedAt = converter.TimeToUnixMilli(source.CreatedAt)
	return typesDictionaryInfo
}
func (c *ConverterImpl) entSysOauthProviderToTypesOauthProviderInfo(source ent.SysOauthProvider) types.OauthProviderInfo {
	var typesOauthProviderInfo types.OauthProviderInfo
	typesOauthProviderInfo.ID = source.ID
	typesOauthProviderInfo.Name = source.Name
	typesOauthProviderInfo.ClientID = source.ClientID
	typesOauthProviderInfo.ClientSecret = source.ClientSecret
	typesOauthProviderInfo.RedirectURL = source.RedirectURL
	typesOauthProviderInfo.Scopes = source.Scopes
	typesOauthProviderInfo.AuthURL = source.AuthURL
	typesOauthProviderInfo.TokenURL = source.TokenURL
	typesOauthProviderInfo.AuthStyle = source.AuthStyle
	typesOauthProviderInfo.InfoURL = source.InfoURL
	typesOauthProviderInfo.CreatedAt = converter.TimeToUnixMilli(source.CreatedAt)
	return typesOauthProviderInfo
}
func (c *ConverterImpl) entSysRoleToTypesRoleInfo(source ent.SysRole) types.RoleInfo {
	var typesRoleInfo types.RoleInfo
	typesRoleInfo.ID = source.ID
	typesRoleInfo.Name = source.Name
	typesRoleInfo.Value = source.Value
	typesRoleInfo.DefaultRouter = source.DefaultRouter
	typesRoleInfo.Status = converter.StatusToUint8(source.Status)
	typesRoleInfo.Remark = source.Remark
	typesRoleInfo.OrderNo = source.OrderNo
	return typesRoleInfo
}
func (c *ConverterImpl) entSysTokenToTypesTokenInfo(source ent.SysToken) types.TokenInfo {
	var typesTokenInfo types.TokenInfo
	typesTokenInfo.ID = source.ID
	typesTokenInfo.UUID = source.UUID
	typesTokenInfo.Token = source.Token
	typesTokenInfo.Source = source.Source
	typesTokenInfo.Status = converter.StatusToUint8(source.Status)
	typesTokenInfo.CreatedAt = converter.TimeToUnixMilli(source.CreatedAt)
	typesTokenInfo.ExpiredAt = converter.TimeToUnixMilli(source.ExpiredAt)
	return typesTokenInfo
}
func (c *ConverterImpl) entSysUserToTypesUserInfo(source ent.SysUser) types.UserInfo {
	var typesUserInfo types.UserInfo
	typesUserInfo.ID = source.ID
	typesUserInfo.UUID = source.UUID
	typesUserInfo.Username = source.Username
	typesUserInfo.Nickname = source.Nickname
	typesUserInfo.Mobile = source.Mobile
	typesUserInfo.RoleID = source.RoleID
	typesUserInfo.Email = source.Email
	typesUserInfo.Avatar = source.Avatar
	typesUserInfo.SideMode = source.SideMode
	typesUserInfo.Status = converter.StatusToUint8(source.Status)
	typesUserInfo.CreatedAt = converter.TimeToUnixMilli(source.CreatedAt)
	typesUserInfo.UpdatedAt = converter.TimeToUnixMilli(source.UpdatedAt)
	return typesUserInfo
}
