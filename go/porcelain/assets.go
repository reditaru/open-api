package porcelain

import (
	"github.com/Sirupsen/logrus"
	"github.com/netlify/open-api/go/models"
	"github.com/netlify/open-api/go/plumbing/operations"
	"github.com/netlify/open-api/go/porcelain/context"
)

func (n *Netlify) AddSiteAsset(ctx context.Context, params *operations.CreateSiteAssetParams) (*models.AssetSignature, error) {
	l := context.GetLogger(ctx)
	l.WithFields(logrus.Fields{
		"site_id": params.SiteID,
	}).Debug("Creating site asset signature")

	resp, err := n.Netlify.Operations.CreateSiteAsset(params, context.GetAuthInfo(ctx))
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (n *Netlify) UpdateSiteAsset(ctx context.Context, params *operations.UpdateSiteAssetParams) (*models.Asset, error) {
	l := context.GetLogger(ctx)
	l.WithFields(logrus.Fields{
		"site_id": params.SiteID,
	}).Debug("Updating site asset state")

	resp, err := n.Netlify.Operations.UpdateSiteAsset(params, context.GetAuthInfo(ctx))
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (n *Netlify) ListSiteAssets(ctx context.Context, params *operations.ListSiteAssetsParams) ([]*models.Asset, error) {
	l := context.GetLogger(ctx)
	l.WithFields(logrus.Fields{
		"site_id": params.SiteID,
	}).Debug("Listing site assets")

	resp, err := n.Netlify.Operations.ListSiteAssets(params, context.GetAuthInfo(ctx))
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (n *Netlify) ShowSiteAssetInfo(ctx context.Context, params *operations.GetSiteAssetInfoParams, showSignature bool) (*models.Asset, error) {
	l := context.GetLogger(ctx)
	l.WithFields(logrus.Fields{
		"site_id":  params.SiteID,
		"asset_id": params.AssetID,
	}).Debug("Show site asset information")

	authInfo := context.GetAuthInfo(ctx)

	resp, err := n.Netlify.Operations.GetSiteAssetInfo(params, authInfo)
	if err != nil {
		return nil, err
	}

	asset := resp.Payload
	if asset.Visibility == "private" && showSignature {
		sigParams := operations.NewGetSiteAssetPublicSignatureParams().WithSiteID(params.SiteID).WithAssetID(params.AssetID)

		sig, err := n.GetSiteAssetPublicSignature(ctx, sigParams)
		if err != nil {
			return nil, err
		}

		asset.URL = sig.URL
	}

	return asset, nil
}

func (n *Netlify) GetSiteAssetPublicSignature(ctx context.Context, params *operations.GetSiteAssetPublicSignatureParams) (*models.AssetPublicSignature, error) {
	l := context.GetLogger(ctx)
	l.WithFields(logrus.Fields{
		"site_id":  params.SiteID,
		"asset_id": params.AssetID,
	}).Debug("Get site asset public signature")

	authInfo := context.GetAuthInfo(ctx)

	resp, err := n.Netlify.Operations.GetSiteAssetPublicSignature(params, authInfo)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
