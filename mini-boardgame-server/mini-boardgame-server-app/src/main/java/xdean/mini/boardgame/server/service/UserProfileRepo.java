package xdean.mini.boardgame.server.service;

import org.springframework.data.jpa.repository.JpaRepository;

import xdean.mini.boardgame.server.model.UserProfile;

public interface UserProfileRepo extends JpaRepository<UserProfile, String> {

}
