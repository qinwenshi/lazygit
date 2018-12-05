package i18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// addDutch will add all dutch translations
func addDutch(i18nObject *i18n.Bundle) error {

	// add the translations
	return i18nObject.AddMessages(language.Dutch,
		&i18n.Message{
			ID:    "NotEnoughSpace",
			Other: "Niet genoeg ruimte om de panelen te renderen",
		}, &i18n.Message{
			ID:    "DiffTitle",
			Other: "Diff",
		}, &i18n.Message{
			ID:    "FilesTitle",
			Other: "Bestanden",
		}, &i18n.Message{
			ID:    "BranchesTitle",
			Other: "Branches",
		}, &i18n.Message{
			ID:    "CommitsTitle",
			Other: "Commits",
		}, &i18n.Message{
			ID:    "StashTitle",
			Other: "Stash",
		}, &i18n.Message{
			ID:    "CommitMessage",
			Other: "Commit bericht",
		}, &i18n.Message{
			ID:    "CommitChanges",
			Other: "Commit veranderingen",
		}, &i18n.Message{
			ID:    "AmendLastCommit",
			Other: "wijzig laatste commit",
		}, &i18n.Message{
			ID:    "SureToAmend",
			Other: "Weet je zeker dat je de laatste commit wilt wijzigen? U kunt het commit-bericht wijzigen vanuit het commits-paneel.",
		}, &i18n.Message{
			ID:    "NoCommitToAmend",
			Other: "Er is geen commits om te wijzigen.",
		}, &i18n.Message{
			ID:    "CommitChangesWithEditor",
			Other: "commit veranderingen met de git editor",
		}, &i18n.Message{
			ID:    "StatusTitle",
			Other: "Status",
		}, &i18n.Message{
			ID:    "GlobalTitle",
			Other: "Global",
		}, &i18n.Message{
			ID:    "navigate",
			Other: "navigeer",
		}, &i18n.Message{
			ID:    "menu",
			Other: "menu",
		}, &i18n.Message{
			ID:    "execute",
			Other: "uitvoeren",
		}, &i18n.Message{
			ID:    "stashFiles",
			Other: "stash-bestanden",
		}, &i18n.Message{
			ID:    "open",
			Other: "open",
		}, &i18n.Message{
			ID:    "ignore",
			Other: "negeren",
		}, &i18n.Message{
			ID:    "delete",
			Other: "verwijderen",
		}, &i18n.Message{
			ID:    "toggleStaged",
			Other: "toggle staged",
		}, &i18n.Message{
			ID:    "toggleStagedAll",
			Other: "toggle staged alle",
		}, &i18n.Message{
			ID:    "refresh",
			Other: "verversen",
		}, &i18n.Message{
			ID:    "push",
			Other: "push",
		}, &i18n.Message{
			ID:    "pull",
			Other: "pull",
		}, &i18n.Message{
			ID:    "addPatch",
			Other: "bewerkingen toevoegen",
		}, &i18n.Message{
			ID:    "edit",
			Other: "bewerken",
		}, &i18n.Message{
			ID:    "scroll",
			Other: "scroll",
		}, &i18n.Message{
			ID:    "abortMerge",
			Other: "samenvoegen afbreken",
		}, &i18n.Message{
			ID:    "resolveMergeConflicts",
			Other: "los merge conflicten op",
		}, &i18n.Message{
			ID:    "checkout",
			Other: "uitchecken",
		}, &i18n.Message{
			ID:    "NoChangedFiles",
			Other: "Geen bestanden veranderd",
		}, &i18n.Message{
			ID:    "FileHasNoUnstagedChanges",
			Other: "Het bestand heeft geen unstaged veranderingen om toe te voegen",
		}, &i18n.Message{
			ID:    "CannotGitAdd",
			Other: "Kan commando niet uitvoeren git add --path untracked files",
		}, &i18n.Message{
			ID:    "CantIgnoreTrackFiles",
			Other: "Kan gevolgde bestanden niet negeren",
		}, &i18n.Message{
			ID:    "NoStagedFilesToCommit",
			Other: "Er zijn geen staged bestanden om te commiten",
		}, &i18n.Message{
			ID:    "NoFilesDisplay",
			Other: "Geen bestanden om te laten zien",
		}, &i18n.Message{
			ID:    "NotAFile",
			Other: "Dit is geen bestand",
		}, &i18n.Message{
			ID:    "PullWait",
			Other: "Pulling...",
		}, &i18n.Message{
			ID:    "PushWait",
			Other: "Pushing...",
		}, &i18n.Message{
			ID:    "FileNoMergeCons",
			Other: "Dit bestand heeft geen merge conflicten",
		}, &i18n.Message{
			ID:    "SureResetHardHead",
			Other: "Weet je het zeker dat je `reset --hard HEAD` wil uitvoeren? Het kan dat je hierdoor bestanden verliest",
		}, &i18n.Message{
			ID:    "SureTo",
			Other: "Weet je het zeker dat je {{.fileName}} wilt {{.deleteVerb}} (je veranderingen zullen worden verwijderd)",
		}, &i18n.Message{
			ID:    "AlreadyCheckedOutBranch",
			Other: "Je hebt deze branch al uitgecheckt",
		}, &i18n.Message{
			ID:    "SureForceCheckout",
			Other: "Weet je zeker dat je het uitchecken wil forceren? Al je lokale verandering zullen worden verwijdert",
		}, &i18n.Message{
			ID:    "ForceCheckoutBranch",
			Other: "Forceer uitchecken op deze branch",
		}, &i18n.Message{
			ID:    "BranchName",
			Other: "Branch naam",
		}, &i18n.Message{
			ID:    "NewBranchNameBranchOff",
			Other: "Nieuw branch naam (Branch is afgeleid van {{.branchName}})",
		}, &i18n.Message{
			ID:    "CantDeleteCheckOutBranch",
			Other: "Je kan een uitgecheckte branch niet verwijderen!",
		}, &i18n.Message{
			ID:    "DeleteBranch",
			Other: "Verwijder branch",
		}, &i18n.Message{
			ID:    "DeleteBranchMessage",
			Other: "Weet je zeker dat je branch {{.selectedBranchName}} wilt verwijderen?",
		}, &i18n.Message{
			ID:    "ForceDeleteBranchMessage",
			Other: "Weet je zeker dat je branch {{.selectedBranchName}} geforceerd wil verwijderen?",
		}, &i18n.Message{
			ID:    "rebaseBranch",
			Other: "rebase branch",
		}, &i18n.Message{
			ID:    "CantRebaseOntoSelf",
			Other: "It is not possible to rebase the branch onto itself!",
		}, &i18n.Message{
			ID:    "CantMergeBranchIntoItself",
			Other: "Je kan niet een branch in zichzelf mergen",
		}, &i18n.Message{
			ID:    "forceCheckout",
			Other: "forceer checkout",
		}, &i18n.Message{
			ID:    "merge",
			Other: "samenvoegen",
		}, &i18n.Message{
			ID:    "checkoutByName",
			Other: "uitchecken bij naam",
		}, &i18n.Message{
			ID:    "newBranch",
			Other: "nieuwe branch",
		}, &i18n.Message{
			ID:    "deleteBranch",
			Other: "verwijder branch",
		}, &i18n.Message{
			ID:    "forceDeleteBranch",
			Other: "verwijder branch (forceer)",
		}, &i18n.Message{
			ID:    "NoBranchesThisRepo",
			Other: "Geen branches voor deze repo",
		}, &i18n.Message{
			ID:    "NoTrackingThisBranch",
			Other: "deze branch wordt niet gevolgd",
		}, &i18n.Message{
			ID:    "CommitWithoutMessageErr",
			Other: "Je kan geen commit maken zonder commit bericht",
		}, &i18n.Message{
			ID:    "CloseConfirm",
			Other: "{{.keyBindClose}}: Sluiten, {{.keyBindConfirm}}: Bevestigen",
		}, &i18n.Message{
			ID:    "close",
			Other: "sluiten",
		}, &i18n.Message{
			ID:    "SureResetThisCommit",
			Other: "Weet je het zeker dat je wil resetten naar deze commit?",
		}, &i18n.Message{
			ID:    "ResetToCommit",
			Other: "Reset Naar Commit",
		}, &i18n.Message{
			ID:    "squashDown",
			Other: "squash beneden",
		}, &i18n.Message{
			ID:    "rename",
			Other: "hernoemen",
		}, &i18n.Message{
			ID:    "resetToThisCommit",
			Other: "reset naar deze commit",
		}, &i18n.Message{
			ID:    "fixupCommit",
			Other: "Fixup commit",
		}, &i18n.Message{
			ID:    "NoCommitsThisBranch",
			Other: "Er zijn geen commits voor deze branch",
		}, &i18n.Message{
			ID:    "OnlySquashTopmostCommit",
			Other: "Kan alleen bovenste commit squashen",
		}, &i18n.Message{
			ID:    "YouNoCommitsToSquash",
			Other: "Je hebt geen commits om mee te squashen",
		}, &i18n.Message{
			ID:    "CantFixupWhileUnstagedChanges",
			Other: "Kan geen Fixup uitvoeren op unstaged veranderingen",
		}, &i18n.Message{
			ID:    "Fixup",
			Other: "Fixup",
		}, &i18n.Message{
			ID:    "SureFixupThisCommit",
			Other: "Weet je zeker dat je fixup wil uitvoeren op deze commit? De commit hieronder zol worden squashed in deze",
		}, &i18n.Message{
			ID:    "OnlyRenameTopCommit",
			Other: "Je kan alleen de bovenste commit hernoemen",
		}, &i18n.Message{
			ID:    "renameCommit",
			Other: "hernoem commit",
		}, &i18n.Message{
			ID:    "renameCommitEditor",
			Other: "rename commit with editor",
		}, &i18n.Message{
			ID:    "PotentialErrInGetselectedCommit",
			Other: "Er is mogelijk een error in getSelected Commit (geen match tussen ui en state)",
		}, &i18n.Message{
			ID:    "NoCommitsThisBranch",
			Other: "Geen commits voor deze branch",
		}, &i18n.Message{
			ID:    "Error",
			Other: "Foutmelding",
		}, &i18n.Message{
			ID:    "resizingPopupPanel",
			Other: "resizen popup paneel",
		}, &i18n.Message{
			ID:    "RunningSubprocess",
			Other: "subprocess lopend",
		}, &i18n.Message{
			ID:    "selectHunk",
			Other: "selecteer stuk",
		}, &i18n.Message{
			ID:    "navigateConflicts",
			Other: "navigeer conflicts",
		}, &i18n.Message{
			ID:    "pickHunk",
			Other: "kies stuk",
		}, &i18n.Message{
			ID:    "pickBothHunks",
			Other: "kies beide stukken",
		}, &i18n.Message{
			ID:    "undo",
			Other: "ongedaan maken",
		}, &i18n.Message{
			ID:    "pop",
			Other: "pop",
		}, &i18n.Message{
			ID:    "drop",
			Other: "drop",
		}, &i18n.Message{
			ID:    "apply",
			Other: "toepassen",
		}, &i18n.Message{
			ID:    "NoStashEntries",
			Other: "Geen stash items",
		}, &i18n.Message{
			ID:    "StashDrop",
			Other: "Stash drop",
		}, &i18n.Message{
			ID:    "SureDropStashEntry",
			Other: "Weet je het zeker dat je deze stash entry wil laten vallen?",
		}, &i18n.Message{
			ID:    "NoStashTo",
			Other: "Geen stash voor {{.method}}",
		}, &i18n.Message{
			ID:    "NoTrackedStagedFilesStash",
			Other: "Je hebt geen tracked/staged bestanden om te laten stashen",
		}, &i18n.Message{
			ID:    "StashChanges",
			Other: "Stash veranderingen",
		}, &i18n.Message{
			ID:    "IssntListOfViews",
			Other: "{{.name}} is niet in de lijst van weergaves",
		}, &i18n.Message{
			ID:    "NoViewMachingNewLineFocusedSwitchStatement",
			Other: "Er machen geen weergave met de newLineFocused switch declaratie",
		}, &i18n.Message{
			ID:    "settingPreviewsViewTo",
			Other: "vorige weergave instellen op: {{.oldViewName}}",
		}, &i18n.Message{
			ID:    "newFocusedViewIs",
			Other: "nieuw gefocussed weergave is {{.newFocusedView}}",
		}, &i18n.Message{
			ID:    "CantCloseConfirmationPrompt",
			Other: "Kon de bevestiging prompt niet sluiten: {{.error}}",
		}, &i18n.Message{
			ID:    "NoChangedFiles",
			Other: "Geen veranderde files",
		}, &i18n.Message{
			ID:    "ClearFilePanel",
			Other: "maak bestandsvenster leeg",
		}, &i18n.Message{
			ID:    "MergeAborted",
			Other: "Merge afgebroken",
		}, &i18n.Message{
			ID:    "OpenConfig",
			Other: "open config file",
		}, &i18n.Message{
			ID:    "EditConfig",
			Other: "verander config file",
		}, &i18n.Message{
			ID:    "ForcePush",
			Other: "Forceer push",
		}, &i18n.Message{
			ID:    "ForcePushPrompt",
			Other: "Jouw branch is afgeweken van de remote branch. Druk 'esc' om te annuleren, of 'enter' om geforceert te pushen.",
		}, &i18n.Message{
			ID:    "checkForUpdate",
			Other: "check voor updates",
		}, &i18n.Message{
			ID:    "CheckingForUpdates",
			Other: "zoeken naar updates...",
		}, &i18n.Message{
			ID:    "OnLatestVersionErr",
			Other: "Je hebt al de laatste versie",
		}, &i18n.Message{
			ID:    "MajorVersionErr",
			Other: "Nieuwe versie ({{.newVersion}}) is niet backwards compatibele vergeleken met de huidige versie ({{.currentVersion}})",
		}, &i18n.Message{
			ID:    "CouldNotFindBinaryErr",
			Other: "Kon geen binary vinden op {{.url}}",
		}, &i18n.Message{
			ID:    "AnonymousReportingTitle",
			Other: "Help lazygit te verbeteren",
		}, &i18n.Message{
			ID:    "AnonymousReportingPrompt",
			Other: "Zou je anonieme data rapportage willen aanzetten om lazygit beter te kunnen maken? (enter/esc)",
		}, &i18n.Message{
			ID:    "GitconfigParseErr",
			Other: `Gogit kon je gitconfig bestand niet goed parsen door de aanwezigheid van losstaande '\' tekens. Het weghalen van deze tekens zou het probleem moeten oplossen. `,
		}, &i18n.Message{
			ID:    "removeFile",
			Other: `Verwijder als untracked / uitchecken wordt gevolgd (ga weg)`,
		}, &i18n.Message{
			ID:    "editFile",
			Other: `verander bestand`,
		}, &i18n.Message{
			ID:    "openFile",
			Other: `open bestand`,
		}, &i18n.Message{
			ID:    "ignoreFile",
			Other: `voeg toe aan .gitignore`,
		}, &i18n.Message{
			ID:    "refreshFiles",
			Other: `refresh bestanden`,
		}, &i18n.Message{
			ID:    "resetHard",
			Other: `harde reset`,
		}, &i18n.Message{
			ID:    "mergeIntoCurrentBranch",
			Other: `merge in met huidige checked out branch`,
		}, &i18n.Message{
			ID:    "ConfirmQuit",
			Other: `Weet je zeker dat je dit programma wil sluiten?`,
		}, &i18n.Message{
			ID:    "SwitchRepo",
			Other: "wissel naar een recente repo",
		}, &i18n.Message{
			ID:    "UnsupportedGitService",
			Other: `Niet-ondersteunde git-service`,
		}, &i18n.Message{
			ID:    "createPullRequest",
			Other: `maak een pull-aanvraag`,
		}, &i18n.Message{
			ID:    "NoBranchOnRemote",
			Other: `Deze branch bestaat niet op de remote. U moet het eerst naar de remote pushen.`,
		}, &i18n.Message{
			ID:    "StageLines",
			Other: `stage individual hunks/lines`,
		}, &i18n.Message{
			ID:    "FileStagingRequirements",
			Other: `Can only stage individual lines for tracked files with unstaged changes`,
		}, &i18n.Message{
			ID:    "StagingTitle",
			Other: `Staging`,
		}, &i18n.Message{
			ID:    "StageHunk",
			Other: `stage hunk`,
		}, &i18n.Message{
			ID:    "StageLine",
			Other: `stage line`,
		}, &i18n.Message{
			ID:    "EscapeStaging",
			Other: `return to files panel`,
		}, &i18n.Message{
			ID:    "CantFindHunks",
			Other: `Could not find any hunks in this patch`,
		}, &i18n.Message{
			ID:    "CantFindHunk",
			Other: `Could not find hunk`,
		},
	)
}
