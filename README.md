# Unity Scene Analysis

Tools to help diagnose file bloat in a unity scene file.

## Installation

```
go install ./cmd/scene-tools
```

## Usage

```bash
scene-tools examine --file Demo.unity --limit 30
```

Example Output:

```
================================================================================
| 26 Categories Found
================================================================================
12243 PrefabInstance 26 mb (89.96%)
12337 Transform 2 mb (7.62%)
6 ParticleSystem 629 kb (2.07%)
107 GameObject 38 kb (0.13%)
21 MeshRenderer 19 kb (0.07%)
18 MonoBehaviour 12 kb (0.04%)
6 ParticleSystemRenderer 8 kb (0.03%)
21 MeshFilter 5 kb (0.02%)
14 BoxCollider 4 kb (0.01%)
2 AudioSource 3 kb (0.01%)
2 Camera 1 kb (0.01%)
6 OcclusionArea 1 kb (0.01%)
1 LightmapSettings 1 kb (0.01%)
1 Light 1 kb (0.00%)
2 RectTransform 1 kb (0.00%)
3 MeshCollider 1 kb (0.00%)
1 RenderSettings 1015 b (0.00%)
1 Terrain 1000 b (0.00%)
1 LODGroup 741 b (0.00%)
1 Canvas 475 b (0.00%)
1 NavMeshSettings 410 b (0.00%)
1 TerrainCollider 308 b (0.00%)
1 OcclusionCullingSettings 283 b (0.00%)
1 Behaviour 241 b (0.00%)
1 CanvasRenderer 185 b (0.00%)
1 AudioListener 173 b (0.00%)
================================================================================
| 24800 Entries Found (30 shown)
================================================================================
ParticleSystem: 4656 lines [316898-321554] 105 kb (0.35%)
ParticleSystem: 4656 lines [221657-226313] 105 kb (0.35%)
ParticleSystem: 4656 lines [306245-310901] 105 kb (0.35%)
ParticleSystem: 4656 lines [394037-398693] 105 kb (0.35%)
ParticleSystem: 4621 lines [504451-509072] 104 kb (0.34%)
ParticleSystem: 4621 lines [187452-192073] 104 kb (0.34%)
Transform: 2820 lines [273597-276417] 64 kb (0.21%)
Transform: 2223 lines [613685-615908] 51 kb (0.17%)
Transform: 2185 lines [777161-779346] 50 kb (0.16%)
Transform: 1755 lines [587927-589682] 40 kb (0.13%)
Transform: 828 lines [410429-411257] 19 kb (0.06%)
Transform: 655 lines [469411-470066] 15 kb (0.05%)
PrefabInstance: 197 lines [314480-314677] 7 kb (0.03%)
        Source Prefab: {fileID: 100100000, guid: 9928698465ca5dc4d930a0b7f31dced9, type: 3}

PrefabInstance: 179 lines [149642-149821] 7 kb (0.02%)
        Source Prefab: {fileID: 100100000, guid: df46d658f4b8a744a962018f4cef4d09, type: 3}

Transform: 317 lines [480139-480456] 7 kb (0.02%)
PrefabInstance: 142 lines [552697-552839] 6 kb (0.02%)
        Source Prefab: {fileID: 100100000, guid: 45f02fac56c35b04cbc6e3f023384597, type: 3}

PrefabInstance: 139 lines [418593-418732] 5 kb (0.02%)
        Source Prefab: {fileID: 100100000, guid: cce5865e2533b764883b42f2466d3e84, type: 3}

Transform: 250 lines [323157-323407] 5 kb (0.02%)
PrefabInstance: 145 lines [235888-236033] 5 kb (0.02%)
        Source Prefab: {fileID: 100100000, guid: 4222d3bd8d11edc4a808cb2a283aa227, type: 3}

PrefabInstance: 131 lines [445216-445347] 5 kb (0.02%)
        Source Prefab: {fileID: 100100000, guid: cce5865e2533b764883b42f2466d3e84, type: 3}

PrefabInstance: 130 lines [573118-573248] 5 kb (0.02%)
        Source Prefab: {fileID: 100100000, guid: 4222d3bd8d11edc4a808cb2a283aa227, type: 3}

PrefabInstance: 126 lines [139569-139695] 5 kb (0.02%)
        Source Prefab: {fileID: 100100000, guid: 4222d3bd8d11edc4a808cb2a283aa227, type: 3}

PrefabInstance: 132 lines [291061-291193] 4 kb (0.02%)
        Source Prefab: {fileID: 100100000, guid: 0ab1a12825c9a1e4cbca267b5e59852e, type: 3}

PrefabInstance: 107 lines [208636-208743] 4 kb (0.02%)
        Source Prefab: {fileID: 100100000, guid: 9064929bb9122724684ed7abbf313f4d, type: 3}

PrefabInstance: 111 lines [13900-14011] 4 kb (0.01%)
        Source Prefab: {fileID: 100100000, guid: 45f02fac56c35b04cbc6e3f023384597, type: 3}

PrefabInstance: 108 lines [483716-483824] 4 kb (0.01%)
        Source Prefab: {fileID: 100100000, guid: 4222d3bd8d11edc4a808cb2a283aa227, type: 3}

PrefabInstance: 99 lines [29076-29175] 4 kb (0.01%)
        Source Prefab: {fileID: 100100000, guid: 0a9df222064d4534da4279115fea376b, type: 3}

PrefabInstance: 95 lines [753956-754051] 4 kb (0.01%)
        Source Prefab: {fileID: 100100000, guid: 0eff1998a022f0a4c8e76aff0ccfdfea, type: 3}

PrefabInstance: 94 lines [765230-765324] 4 kb (0.01%)
        Source Prefab: {fileID: 100100000, guid: 7ae254bddcf0def478298b88c82e882a, type: 3}

PrefabInstance: 96 lines [341832-341928] 3 kb (0.01%)
        Source Prefab: {fileID: 100100000, guid: f38205ce7c90f6b44854e3ce06c8f473, type: 3}
```