package main

import (
    "errors"
    "net"
    "net/http"
    "net/url"
)

func validatePublicHTTPURL(raw string) error {
    u, err := url.Parse(raw)
    if err != nil || (u.Scheme != "https" && u.Scheme != "http") {
        return errors.New("invalid url")
    }
    ips, err := net.LookupIP(u.Hostname())
    if err != nil {
        return err
    }
    for _, ip := range ips {
        if ip.IsPrivate() || ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsUnspecified() {
            return errors.New("blocked address")
        }
    }
    return nil
}

func fetchPeerInfo(peerURL string) error {
    err := validatePublicHTTPURL(peerURL)
    if err != nil {
        return err
    }
    resp, err := http.Get(peerURL)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    return nil
}
