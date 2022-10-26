// Code generated by ent, DO NOT EDIT.

package ent

type JsonbMode string

const (
	JsonbReplace JsonbMode = "replace"
	JsonbMerge   JsonbMode = "merge"
)

func (sauo *SysApiUpdateOne) Copy(input *UpdateSysApiInput, jsonbMode JsonbMode) *SysApiUpdateOne {

	if input.Path != nil {

		sauo.SetPath(*input.Path)

	}

	if input.Description != nil {

		sauo.SetDescription(*input.Description)

	}

	if input.APIGroup != nil {

		sauo.SetAPIGroup(*input.APIGroup)

	}

	if input.Method != nil {

		sauo.SetMethod(*input.Method)

	}

	if input.CreatedAt != nil {

		sauo.SetCreatedAt(*input.CreatedAt)

	}

	if input.UpdatedAt != nil {

		sauo.SetUpdatedAt(*input.UpdatedAt)

	}

	if input.DeletedAt != nil {

		sauo.SetDeletedAt(*input.DeletedAt)

	}
	return sauo
}

func (sduo *SysDictionaryUpdateOne) Copy(input *UpdateSysDictionaryInput, jsonbMode JsonbMode) *SysDictionaryUpdateOne {

	if input.Title != nil {

		sduo.SetTitle(*input.Title)

	}

	if input.Name != nil {

		sduo.SetName(*input.Name)

	}

	if input.Status != nil {

		sduo.SetStatus(*input.Status)

	}

	if input.Desc != nil {

		sduo.SetDesc(*input.Desc)

	}

	if input.CreatedAt != nil {

		sduo.SetCreatedAt(*input.CreatedAt)

	}

	if input.UpdatedAt != nil {

		sduo.SetUpdatedAt(*input.UpdatedAt)

	}

	if input.DeletedAt != nil {

		sduo.SetDeletedAt(*input.DeletedAt)

	}
	return sduo
}

func (sdduo *SysDictionaryDetailUpdateOne) Copy(input *UpdateSysDictionaryDetailInput, jsonbMode JsonbMode) *SysDictionaryDetailUpdateOne {

	if input.Title != nil {

		sdduo.SetTitle(*input.Title)

	}

	if input.Key != nil {

		sdduo.SetKey(*input.Key)

	}

	if input.Value != nil {

		sdduo.SetValue(*input.Value)

	}

	if input.Status != nil {

		sdduo.SetStatus(*input.Status)

	}

	if input.Remark != nil {

		sdduo.SetRemark(*input.Remark)

	}

	if input.OrderNo != nil {

		sdduo.SetOrderNo(*input.OrderNo)

	}

	if input.CreatedAt != nil {

		sdduo.SetCreatedAt(*input.CreatedAt)

	}

	if input.UpdatedAt != nil {

		sdduo.SetUpdatedAt(*input.UpdatedAt)

	}

	if input.DeletedAt != nil {

		sdduo.SetDeletedAt(*input.DeletedAt)

	}
	return sdduo
}

func (smuo *SysMenuUpdateOne) Copy(input *UpdateSysMenuInput, jsonbMode JsonbMode) *SysMenuUpdateOne {

	if input.MenuLevel != nil {

		smuo.SetMenuLevel(*input.MenuLevel)

	}

	if input.MenuType != nil {

		smuo.SetMenuType(*input.MenuType)

	}

	if input.ParentID != nil {

		smuo.SetParentID(*input.ParentID)

	}

	if input.Path != nil {

		smuo.SetPath(*input.Path)

	}

	if input.Name != nil {

		smuo.SetName(*input.Name)

	}

	if input.Redirect != nil {

		smuo.SetRedirect(*input.Redirect)

	}

	if input.Component != nil {

		smuo.SetComponent(*input.Component)

	}

	if input.OrderNo != nil {

		smuo.SetOrderNo(*input.OrderNo)

	}

	if input.Disabled != nil {

		smuo.SetDisabled(*input.Disabled)

	}

	if input.Meta != nil {

		smuo.SetMeta(*input.Meta)

	}

	if input.CreatedAt != nil {

		smuo.SetCreatedAt(*input.CreatedAt)

	}

	if input.UpdatedAt != nil {

		smuo.SetUpdatedAt(*input.UpdatedAt)

	}

	if input.DeletedAt != nil {

		smuo.SetDeletedAt(*input.DeletedAt)

	}
	return smuo
}

func (smpuo *SysMenuParamUpdateOne) Copy(input *UpdateSysMenuParamInput, jsonbMode JsonbMode) *SysMenuParamUpdateOne {

	if input.MenuID != nil {

		smpuo.SetMenuID(*input.MenuID)

	}

	if input.Type != nil {

		smpuo.SetType(*input.Type)

	}

	if input.Key != nil {

		smpuo.SetKey(*input.Key)

	}

	if input.Value != nil {

		smpuo.SetValue(*input.Value)

	}

	if input.CreatedAt != nil {

		smpuo.SetCreatedAt(*input.CreatedAt)

	}

	if input.UpdatedAt != nil {

		smpuo.SetUpdatedAt(*input.UpdatedAt)

	}

	if input.DeletedAt != nil {

		smpuo.SetDeletedAt(*input.DeletedAt)

	}
	return smpuo
}

func (sopuo *SysOauthProviderUpdateOne) Copy(input *UpdateSysOauthProviderInput, jsonbMode JsonbMode) *SysOauthProviderUpdateOne {

	if input.Name != nil {

		sopuo.SetName(*input.Name)

	}

	if input.ClientID != nil {

		sopuo.SetClientID(*input.ClientID)

	}

	if input.ClientSecret != nil {

		sopuo.SetClientSecret(*input.ClientSecret)

	}

	if input.RedirectURL != nil {

		sopuo.SetRedirectURL(*input.RedirectURL)

	}

	if input.Scopes != nil {

		sopuo.SetScopes(*input.Scopes)

	}

	if input.AuthURL != nil {

		sopuo.SetAuthURL(*input.AuthURL)

	}

	if input.TokenURL != nil {

		sopuo.SetTokenURL(*input.TokenURL)

	}

	if input.AuthStyle != nil {

		sopuo.SetAuthStyle(*input.AuthStyle)

	}

	if input.InfoURL != nil {

		sopuo.SetInfoURL(*input.InfoURL)

	}

	if input.CreatedAt != nil {

		sopuo.SetCreatedAt(*input.CreatedAt)

	}

	if input.UpdatedAt != nil {

		sopuo.SetUpdatedAt(*input.UpdatedAt)

	}

	if input.DeletedAt != nil {

		sopuo.SetDeletedAt(*input.DeletedAt)

	}
	return sopuo
}

func (sruo *SysRoleUpdateOne) Copy(input *UpdateSysRoleInput, jsonbMode JsonbMode) *SysRoleUpdateOne {

	if input.Name != nil {

		sruo.SetName(*input.Name)

	}

	if input.Value != nil {

		sruo.SetValue(*input.Value)

	}

	if input.DefaultRouter != nil {

		sruo.SetDefaultRouter(*input.DefaultRouter)

	}

	if input.Status != nil {

		sruo.SetStatus(*input.Status)

	}

	if input.Remark != nil {

		sruo.SetRemark(*input.Remark)

	}

	if input.OrderNo != nil {

		sruo.SetOrderNo(*input.OrderNo)

	}

	if input.CreatedAt != nil {

		sruo.SetCreatedAt(*input.CreatedAt)

	}

	if input.UpdatedAt != nil {

		sruo.SetUpdatedAt(*input.UpdatedAt)

	}

	if input.DeletedAt != nil {

		sruo.SetDeletedAt(*input.DeletedAt)

	}
	return sruo
}

func (stuo *SysTokenUpdateOne) Copy(input *UpdateSysTokenInput, jsonbMode JsonbMode) *SysTokenUpdateOne {

	if input.UUID != nil {

		stuo.SetUUID(*input.UUID)

	}

	if input.Token != nil {

		stuo.SetToken(*input.Token)

	}

	if input.Source != nil {

		stuo.SetSource(*input.Source)

	}

	if input.Status != nil {

		stuo.SetStatus(*input.Status)

	}

	if input.ExpiredAt != nil {

		stuo.SetExpiredAt(*input.ExpiredAt)

	}

	if input.CreatedAt != nil {

		stuo.SetCreatedAt(*input.CreatedAt)

	}

	if input.UpdatedAt != nil {

		stuo.SetUpdatedAt(*input.UpdatedAt)

	}

	if input.DeletedAt != nil {

		stuo.SetDeletedAt(*input.DeletedAt)

	}
	return stuo
}

func (suuo *SysUserUpdateOne) Copy(input *UpdateSysUserInput, jsonbMode JsonbMode) *SysUserUpdateOne {

	if input.UUID != nil {

		suuo.SetUUID(*input.UUID)

	}

	if input.Username != nil {

		suuo.SetUsername(*input.Username)

	}

	if input.Password != nil {

		suuo.SetPassword(*input.Password)

	}

	if input.Nickname != nil {

		suuo.SetNickname(*input.Nickname)

	}

	if input.SideMode != nil {

		suuo.SetSideMode(*input.SideMode)

	}

	if input.Avatar != nil {

		suuo.SetAvatar(*input.Avatar)

	}

	if input.BaseColor != nil {

		suuo.SetBaseColor(*input.BaseColor)

	}

	if input.ActiveColor != nil {

		suuo.SetActiveColor(*input.ActiveColor)

	}

	if input.RoleID != nil {

		suuo.SetRoleID(*input.RoleID)

	}

	if input.Mobile != nil {

		suuo.SetMobile(*input.Mobile)

	}

	if input.Email != nil {

		suuo.SetEmail(*input.Email)

	}

	if input.Status != nil {

		suuo.SetStatus(*input.Status)

	}

	if input.CreatedAt != nil {

		suuo.SetCreatedAt(*input.CreatedAt)

	}

	if input.UpdatedAt != nil {

		suuo.SetUpdatedAt(*input.UpdatedAt)

	}

	if input.DeletedAt != nil {

		suuo.SetDeletedAt(*input.DeletedAt)

	}
	return suuo
}