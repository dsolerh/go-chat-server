package main

import (
	"testing"

	"github.com/stretchr/objx"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)

	url, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("AuthAvatar.GetAvatarURL should return ErrNoAvatarURL when no value present")
	}

	// set a value
	testURL := "http://url-to-gravatar/"
	client.userData = objx.Map{"avatar_url": testURL}

	url, err = authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("AuthAvatar.GetAvatarURL should return no error when value present")
	}

	if url != testURL {
		t.Error("AuthAvatar.GetAvatarURL should return the correct URL")
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	client := new(client)
	client.userData = objx.Map{
		"userid": "55502f40dc8b7c769880b10874abc9d0",
	}
	url, err := gravatarAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("GravatarAvatar.GetAvatarURL should not return an error")
	}
	if url != "//www.gravatar.com/avatar/55502f40dc8b7c769880b10874abc9d0" {
		t.Errorf("GravatarAvatar.GetAvatarURL wrongly returned %s", url)
	}
}
